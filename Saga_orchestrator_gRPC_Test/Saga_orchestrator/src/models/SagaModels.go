package models

type SagaStep struct {
	Name             string
	Execute          string
	Compensate       string
	ExecuteParams    []string
	CompensateParams []string
}

type SagaState string

const (
	NotStarted  SagaState = "NotStarted"
	InProgress  SagaState = "InProgress"
	Completed   SagaState = "Completed"
	Failed      SagaState = "Failed"
	CustomState SagaState = "CustomState"
)

type SagaInstance struct {
	ID          string
	CurrentStep int
	State       SagaState
	Steps       []SagaStep
}
