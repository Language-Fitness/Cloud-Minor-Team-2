package Saga

type SagaStep struct {
	Name       string
	Execute    func() error
	Compensate func() error
}

type SagaInstance struct {
	ID          string
	CurrentStep int
	Steps       []SagaStep
}
