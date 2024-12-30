package saga

import "time"

// SagaMode defines the execution mode of the Saga (synchronous or asynchronous).
type SagaMode string

const (
	// SyncMode represents sequential execution of Saga steps.
	SyncMode SagaMode = "sync"
	// AsyncMode represents concurrent execution of Saga steps.
	AsyncMode SagaMode = "async"
)

// Saga defines the interface for managing a Saga workflow.
type Saga interface {
	AddStep(action func() (interface{}, error), compensation func(interface{}) error)
	Execute() error
	SetTimeout(timeout time.Duration)
	SetMode(mode SagaMode)
	EnableExecutionTimeLogging(enable bool)
}

// SagaStep represents a single step in a Saga workflow.
type SagaStep struct {
	// Action defines the primary action for the Saga step.
	// It returns a response and an error if the action fails.
	Action func() (interface{}, error)

	// Compensation defines the rollback mechanism for the Saga step.
	// It takes the response from the Action as input.
	Compensation func(interface{}) error
}
