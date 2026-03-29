package main

import (
	cache_service "github.com/coco1660/cache2go/internal/controller/rpc/idl/kitex_gen/cache_service/cacheservice"
	"log"
)

func main() {
	svr := cache_service.NewServer(new(CacheServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
