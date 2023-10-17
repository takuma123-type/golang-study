package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/golang-study/src/infra/router"
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
