package v1

import (
	"fmt"
	cache2 "github.com/coco1660/cache2go/internal/cache"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetKey(c *gin.Context) {
	cache := c.Param("cache")
	if cache == "" {
		fmt.Errorf("Cache is empty")
		errorResponse(c, http.StatusBadRequest, "Cache is empty")
		return
	}

	key := c.Param("key")
	if key == "" {
		fmt.Errorf("Key is empty")
		errorResponse(c, http.StatusBadRequest, "Key is empty")
		return
	}

	item, err := cache2.Cache(cache).Value(key)
	if err != nil {
		fmt.Errorf("GetKey fail, cache: %s, key: %s, err: %w", cache, key, err)
		errorResponse(c, http.StatusBadRequest, "GetKey fail")
		return
	}
	c.JSON(http.StatusOK, item)
}

func SetKey(c *gin.Context) {
	cache := c.Param("cache")
	if cache == "" {
		fmt.Errorf("Cache is empty")
		errorResponse(c, http.StatusBadRequest, "Cache is empty")
		return
	}

	key := c.Param("key")
	if key == "" {
		fmt.Errorf("Key is empty")
		errorResponse(c, http.StatusBadRequest, "Key is empty")
		return
	}

	value := c.Param("value")
	if value == "" {
		fmt.Errorf("Value is empty")
		errorResponse(c, http.StatusBadRequest, "Value is empty")
		return
	}

	life := c.Param("life")
	var life_span int64
	if life == "" {
		life_span = 0
	} else {
		life_span, _ = strconv.ParseInt(life, 10, 64)
	}

	item := cache2.Cache(cache).Add(key, time.Duration(life_span), value)
	c.JSON(http.StatusOK, item)
}

func DeleteKey(c *gin.Context) {
	cache := c.Param("cache")
	if cache == "" {
		fmt.Errorf("Cache is empty")
		errorResponse(c, http.StatusBadRequest, "Cache is empty")
		return
	}

	key := c.Param("key")
	if key == "" {
		fmt.Errorf("Key is empty")
		errorResponse(c, http.StatusBadRequest, "Key is empty")
		return
	}

	item, err := cache2.Cache(cache).Delete(key)
	if err != nil {
		fmt.Errorf("Delete fail, cache: %s, key: %s, err: %w", cache, key, err)
		errorResponse(c, http.StatusBadRequest, "Key is empty")
		return
	}

	c.JSON(http.StatusOK, item)
}

func Exists(c *gin.Context) {
	cache := c.Param("cache")
	if cache == "" {
		fmt.Errorf("Cache is empty")
		errorResponse(c, http.StatusBadRequest, "Cache is empty")
		return
	}

	key := c.Param("key")
	if key == "" {
		fmt.Errorf("Key is empty")
		errorResponse(c, http.StatusBadRequest, "Key is empty")
		return
	}

	exists := cache2.Cache(cache).Exists(key)

	c.JSON(http.StatusOK, exists)
}
