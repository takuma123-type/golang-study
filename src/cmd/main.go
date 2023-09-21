package main

import (
	"github.com/hrs-o/docker-go/infra/router"
)

func main() {
	r := router.InitRoutes()
	r.Run(":9090")
}
