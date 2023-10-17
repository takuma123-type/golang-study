package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/golang-study/src/infra/middleware"
	"github.com/takuma123-type/golang-study/src/infra/router"
)

func main() {
	g := gin.Default()

	g.Use(middleware.HandleErrorMiddleware())

	router.NewUserRouter(g)

	log.Fatal(g.Run(":8000"))
}
