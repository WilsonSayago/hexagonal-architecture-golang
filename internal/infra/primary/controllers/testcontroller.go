package controllers

import (
	"github.com/gin-gonic/gin"
	"linkedin/internal/core"
	"sync"
)

var (
	testInstance *TestController
	testOnce     sync.Once
)

type TestController struct {
	service core.ITestUseCase
}

func NewTestController(service core.ITestUseCase) *TestController {
	testOnce.Do(func() {
		testInstance = &TestController{
			service: service,
		}
	})
	return testInstance
}

func (o *TestController) RunController(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		response, _ := o.service.GetTest()
		
		c.JSON(200, response)
	})
}
