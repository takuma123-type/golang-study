package main

import (
	"github.com/hrs-o/docker-go/infra/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, Worldaaaaa")
}

func main() {
	r := router.InitRoutes()

	// 新しいエンドポイントの追加
	r.GET("/", helloHandler)

	r.Run(":9090")
}
