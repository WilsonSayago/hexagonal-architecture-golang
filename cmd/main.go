package main

import (
	"github.com/gin-gonic/gin"
	"linkedin/internal/infra/config/instances"
	"linkedin/internal/infra/primary/controllers"
)

func main() {
	
	r := gin.Default()
	
	controllers.NewTestController(instances.GetTestService()).RunController(r)
	
	r.Run()
}
