package repo

import (
	"fmt"
	"github.com/coco1660/cache2go/internal/entity"
	"github.com/coco1660/cache2go/pkg/mysql"
)

type CacheRepo struct {
	*mysql.Mysql
}

// New -.
func New(mysql *mysql.Mysql) *CacheRepo {
	return &CacheRepo{mysql}
}

func (r *CacheRepo) Load() (map[string][]entity.CacheItems, error) {
	var (
		err        error
		tables     []entity.Tables
		items      []entity.CacheItems
		idName     = make(map[int64]string)
		cacheTable = make(map[string][]entity.CacheItems)
		tx         = r.Engine
	)
	// tx.Get操作是只查找一条记录，find是查找所有数据
	if err = tx.Find(&tables); err != nil {
		return nil, fmt.Errorf("Load - GetTables - Fail : %w", err)
	}
	for _, cache := range tables {
		idName[cache.Id] = cache.Name
		cacheTable[cache.Name] = []entity.CacheItems{}
	}
	if err = tx.Find(&items); err != nil {
		return nil, fmt.Errorf("Load - GetItems - Fail : %w", err)
	}
	for _, item := range items {
		name := idName[item.TableID]
		cacheTable[name] = append(cacheTable[name], item)
	}
	return cacheTable, nil
}
func (r *CacheRepo) SaveCache(caches []*entity.Tables) (map[string]int64, error) {
	var (
		err     error
		cacheId = make(map[string]int64)
		tx      = r.Engine
	)

	// 1、清空tables表
	_, err = tx.Exec(fmt.Sprintf("TRUNCATE TABLE %s", "tables"))
	if err != nil {
		return nil, fmt.Errorf("Clear tables - Fail : %w", err)
	}

	for _, cache := range caches {
		id, err := tx.Insert(cache)
		if err != nil {
			return nil, fmt.Errorf("insert cache: %s - Fail : %w", cache.Name, err)
		}
		cacheId[cache.Name] = id
	}
	return cacheId, nil
}
func (r *CacheRepo) SaveItems(cache map[string]*entity.CacheItems) error {
	var (
		err error
		tx  = r.Engine
	)

	// 2、清空CacheItems表
	_, err = tx.Exec(fmt.Sprintf("TRUNCATE TABLE %s", "cache_items"))
	if err != nil {
		fmt.Errorf("Clear cache_items - Fail : %w", err)
	}

	// 4、遍历cache存所有items
	for k, v := range cache {
		_, err = tx.Insert(v)
		if err != nil {
			return fmt.Errorf("insert item fail, cache: %s, key: %s, err: %w", k, v, err)
		}
	}
	return err
}
