package services

import (
	"linkedin/internal/core"
	"sync"
)

var (
	serviceInstance     core.ITestUseCase
	serviceInstanceOnce sync.Once
)

type TestService struct{}

func NewTestService() core.ITestUseCase {
	serviceInstanceOnce.Do(func() {
		service := &TestService{}
		var IService core.ITestUseCase = service
		serviceInstance = IService
	})
	return serviceInstance
}

func (t *TestService) GetTest() (string, error) {
	return "Test", nil
}
