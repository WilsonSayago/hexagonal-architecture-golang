package core

type ITestUseCase interface {
	GetTest() (string, error)
}
