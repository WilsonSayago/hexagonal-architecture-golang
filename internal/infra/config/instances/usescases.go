package instances

import (
	"linkedin/internal/core"
	"linkedin/internal/core/services"
)

func GetTestService() core.ITestUseCase {
	return services.NewTestService()
}
