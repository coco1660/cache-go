package app

import (
	"fmt"
	"github.com/coco1660/cache2go/config"
	"github.com/coco1660/cache2go/internal/cache"
	"github.com/coco1660/cache2go/internal/controller/http/v1"
	httpserver2 "github.com/coco1660/cache2go/pkg/httpserver"
	"github.com/coco1660/cache2go/pkg/logger"
	"github.com/coco1660/cache2go/pkg/mysql"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {

	l := logger.New(cfg.Log.Level)

	mysql, err := mysql.New(
		cfg.MySQL.URL,
		mysql.MaxIdleConns(cfg.MySQL.MaxIdleConns),
		mysql.MaxOpenConns(cfg.MySQL.MaxOpenConns),
	)
	if err != nil {
		_ = fmt.Errorf("app - Run - mysql.New: %w", err)
	}
	defer mysql.Close()

	// 服务启动前load
	err = cache.CacheLoad(mysql)
	if err != nil {
		fmt.Errorf("app - Run - cache.CacheLoad: %w", err)
	}

	// HTTP Server
	handler := gin.New()

	v1.NewRouter(handler, l)
	httpServer := httpserver2.New(handler, httpserver2.Port(cfg.HTTP.Port))

	//// gRPC Server
	//grpcServer := grpcserver2.New(grpcserver2.Port(cfg.GRPC.Port))
	//rpc.NewRouter(l)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		cache.CacheSave(mysql, l)
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		cache.CacheSave(mysql, l)
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		//case err = <-grpcServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	//grpcServer.Shutdown()

}
