package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hrs-o/docker-go/interface/controllers"
)

func HealthRoutes(route *gin.Engine) {
	route.GET("/health", controllers.HealthCheck)
}
