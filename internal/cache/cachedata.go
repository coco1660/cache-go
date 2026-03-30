/*
 * Simple caching library with expiration capabilities
 *     Copyright (c) 2012, Radu Ioan Fericean
 *                   2013-2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE.txt
 */

package cache

import (
	"fmt"
	"github.com/coco1660/cache2go/internal/cache/repo"
	"github.com/coco1660/cache2go/internal/entity"
	"github.com/coco1660/cache2go/pkg/logger"
	"github.com/coco1660/cache2go/pkg/mysql"
	"sync"
	"time"
)

var (
	cache = make(map[string]*CacheTable)
	mutex sync.RWMutex
)

func CacheLoad(mysql *mysql.Mysql) error {
	cacheRepo := repo.New(mysql)
	// 需要过滤过期key
	load, err := cacheRepo.Load()
	if err != nil {
		return fmt.Errorf("load failed: %w", err)
	}
	for n, _ := range load {
		cache[n] = &CacheTable{}
	}
	now := time.Now()
	for n, i := range load {
		for _, item := range i {
			if now.Sub(item.ExpireAt) >= 0 {
				continue
			}
			// 还没过期
			cache[n].items[item.Key] = &CacheItem{
				key:         item.Key,
				data:        item.Value,
				lifeSpan:    item.ExpireAt.Sub(now),
				createdOn:   item.CreateTime,
				accessedOn:  item.UpdateTime,
				accessCount: item.AccessCount,
			}
		}
	}
	return nil
}

func CacheSave(mysql *mysql.Mysql, l *logger.Logger) error {
	l.Error(fmt.Errorf("cache save"))
	cacheRepo := repo.New(mysql)
	data := make(map[string]*entity.CacheItems)
	caches := []*entity.Tables{}
	for k, _ := range cache {
		caches = append(caches, &entity.Tables{Name: k})
	}
	cacheId, err := cacheRepo.SaveCache(caches)

	for k, v := range cache {
		l.Error(fmt.Errorf("k: %s", k))
		l.Error(fmt.Errorf("v: %v", v))
		l.Error(fmt.Errorf("v: %v", v.items))
		items := v.items
		for key, value := range items {
			data[key.(string)] = &entity.CacheItems{
				TableID:     cacheId[k],
				Key:         key.(string),
				Value:       value.data.(string),
				ExpireAt:    value.accessedOn.Add(v.items[k].lifeSpan),
				CreateTime:  value.createdOn,
				UpdateTime:  value.accessedOn,
				AccessCount: value.accessCount,
			}
		}
		err = cacheRepo.SaveItems(data)
	}
	l.Error(fmt.Errorf("cache save finished"))
	return err
}

// Cache returns the existing cache table with given name or creates a new one
// if the table does not exist yet.
func Cache(table string) *CacheTable {
	// 获取当前访问的table，加读锁
	mutex.RLock()
	t, ok := cache[table]
	mutex.RUnlock()
	// 不存在，需要创建，加写锁
	if !ok {
		mutex.Lock()
		// 加锁后再次获取，double-check
		t, ok = cache[table]
		// Double check whether the table exists or not.
		if !ok {
			t = &CacheTable{
				name:  table,
				items: make(map[interface{}]*CacheItem),
			}
			cache[table] = t
		}
		mutex.Unlock()
	}
	// 存在直接返回
	return t
}
