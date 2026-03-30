package v1

import (
	"github.com/coco1660/cache2go/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter(handler *gin.Engine, l logger.Interface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers

	h := handler.Group("/v1")
	{
		h.GET("get/:cache/:key", GetKey)
		h.GET("set/:cache/:key/:value/:life", SetKey)
		h.POST("delete/:cache/:key", DeleteKey)
		h.GET("exists/:cache/:key", Exists)
	}
}
