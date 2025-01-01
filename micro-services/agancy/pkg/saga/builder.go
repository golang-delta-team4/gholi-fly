package saga

import "time"

// SagaBuilder helps construct a SagaOrchestrator with a fluent API.
type SagaBuilder struct {
	saga *sagaOrchestrator
}

// NewSagaBuilder initializes a new SagaBuilder.
func NewSagaBuilder() *SagaBuilder {
	return &SagaBuilder{saga: newSagaOrchestrator()}
}

// SetMode sets the execution mode (SyncMode or AsyncMode).
func (b *SagaBuilder) SetMode(mode SagaMode) *SagaBuilder {
	b.saga.SetMode(mode)
	return b
}

// AddStep adds a step to the Saga.
func (b *SagaBuilder) AddStep(action func() (interface{}, error), compensation func(interface{}) error) *SagaBuilder {
	b.saga.AddStep(action, compensation)
	return b
}

// SetTimeout sets a timeout for the Saga.
func (b *SagaBuilder) SetTimeout(timeout time.Duration) *SagaBuilder {
	b.saga.SetTimeout(timeout)
	return b
}

// EnableExecutionTimeLogging enables or disables logging of execution time.
func (b *SagaBuilder) EnableExecutionTimeLogging(enable bool) *SagaBuilder {
	b.saga.EnableExecutionTimeLogging(enable)
	return b
}

// Build returns the constructed SagaOrchestrator.
func (b *SagaBuilder) Build() *sagaOrchestrator {
	return b.saga
}
