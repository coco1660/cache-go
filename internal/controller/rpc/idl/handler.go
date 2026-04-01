package main

import (
	"context"
	"fmt"
	"github.com/coco1660/cache2go/internal/cache"
	"github.com/coco1660/cache2go/internal/controller/rpc/idl/kitex_gen/cache_service"
	"time"
)

// CacheServiceImpl implements the last service interface defined in the IDL.
type CacheServiceImpl struct{}

// Get implements the CacheServiceImpl interface.
func (s *CacheServiceImpl) Get(ctx context.Context, req *cache_service.GetRequest) (resp *cache_service.GetResponse, err error) {
	cacheName := req.GetCache()
	if cacheName == "" {
		return nil, fmt.Errorf("cache name is empty")
	}

	key := req.GetKey()
	if key == "" {
		return nil, fmt.Errorf("cache key is empty")
	}

	value, err := cache.Cache(cacheName).Value(key)
	resp = &cache_service.GetResponse{
		Value: value.Data().(string),
		Base: &cache_service.BaseResp{
			Code: 200,
			Msg:  "success",
		},
	}
	return
}

// Set implements the CacheServiceImpl interface.
func (s *CacheServiceImpl) Set(ctx context.Context, req *cache_service.SetRequest) (resp *cache_service.SetResponse, err error) {
	cacheName := req.GetCache()
	if cacheName == "" {
		return nil, fmt.Errorf("cache name is empty")
	}

	key := req.GetValue()
	if key == "" {
		return nil, fmt.Errorf("cache key is empty")
	}

	value := req.GetValue()
	if value == "" {
		return nil, fmt.Errorf("cache value is empty")
	}

	life := req.GetLifeSpan()

	_ = cache.Cache(cacheName).Add(key, time.Duration(life), value)

	resp = &cache_service.SetResponse{
		Base: &cache_service.BaseResp{
			Code: 200,
			Msg:  "success",
		},
	}
	return
}

// Delete implements the CacheServiceImpl interface.
func (s *CacheServiceImpl) Delete(ctx context.Context, req *cache_service.DeleteRequest) (resp *cache_service.DeleteResponse, err error) {
	cacheName := req.GetCache()
	if cacheName == "" {
		return nil, fmt.Errorf("cache name is empty")
	}

	key := req.GetKey()
	if key == "" {
		return nil, fmt.Errorf("cache key is empty")
	}

	_, err = cache.Cache(cacheName).Delete(key)
	if err != nil {
		resp = &cache_service.DeleteResponse{
			Deleted: false,
			Base: &cache_service.BaseResp{
				Code: 500,
				Msg:  "fail",
			},
		}
	}

	resp = &cache_service.DeleteResponse{
		Deleted: true,
		Base: &cache_service.BaseResp{
			Code: 200,
			Msg:  "success",
		},
	}
	return
}

// Exists implements the CacheServiceImpl interface.
func (s *CacheServiceImpl) Exists(ctx context.Context, req *cache_service.ExistsRequest) (resp *cache_service.ExistsResponse, err error) {
	cacheName := req.GetCache()
	if cacheName == "" {
		return nil, fmt.Errorf("cache name is empty")
	}

	key := req.GetKey()
	if key == "" {
		return nil, fmt.Errorf("cache key is empty")
	}

	res := cache.Cache(cacheName).Exists(key)

	resp = &cache_service.ExistsResponse{
		Exists: res,
		Base: &cache_service.BaseResp{
			Code: 200,
			Msg:  "success",
		},
	}
	return
}

// New_ implements the CacheServiceImpl interface.
func (s *CacheServiceImpl) New_(ctx context.Context, req *cache_service.NewCacheRequest_) (resp *cache_service.NewCacheResponse_, err error) {
	name := req.GetName()

	myCache := cache.Cache(name)

	// This callback will be triggered every time a new item
	// gets added to the cache.
	myCache.SetAddedItemCallback(func(entry *cache.CacheItem) {
		fmt.Println("Added Callback 1:", entry.Key(), entry.Data(), entry.CreatedOn())
	})
	// This callback will be triggered every time an item
	// is about to be removed from the cache.
	myCache.SetAboutToDeleteItemCallback(func(entry *cache.CacheItem) {
		fmt.Println("Deleting:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	return
}
