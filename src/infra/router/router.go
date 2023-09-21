package router

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	r := gin.Default()
	HealthRoutes(r)
	return r
}
