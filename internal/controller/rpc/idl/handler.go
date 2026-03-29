package main

import (
	"context"
	"fmt"
	"github.com/coco1660/cache2go/internal/cache"
	cache_service "github.com/coco1660/cache2go/internal/controller/rpc/idl/kitex_gen/cache_service"
)

// CacheServiceImpl implements the last service interface defined in the IDL.
type CacheServiceImpl struct{}

// Get implements the CacheServiceImpl interface.
func (s *CacheServiceImpl) Get(ctx context.Context, req *cache_service.GetRequest) (resp *cache_service.GetResponse, err error) {

	return
}

// Set implements the CacheServiceImpl interface.
func (s *CacheServiceImpl) Set(ctx context.Context, req *cache_service.SetRequest) (resp *cache_service.SetResponse, err error) {
	// TODO: Your code here...
	return
}

// Delete implements the CacheServiceImpl interface.
func (s *CacheServiceImpl) Delete(ctx context.Context, req *cache_service.DeleteRequest) (resp *cache_service.DeleteResponse, err error) {
	// TODO: Your code here...
	return
}

// Exists implements the CacheServiceImpl interface.
func (s *CacheServiceImpl) Exists(ctx context.Context, req *cache_service.ExistsRequest) (resp *cache_service.ExistsResponse, err error) {
	// TODO: Your code here...
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
