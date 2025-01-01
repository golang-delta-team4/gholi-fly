package saga

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// sagaOrchestrator implements the Saga interface.
type sagaOrchestrator struct {
	mode                SagaMode
	steps               []SagaStep
	errCh               chan error
	context             context.Context
	cancel              context.CancelFunc
	enableExecutionTime bool
	ExecutionTime       time.Duration
}

// newSagaOrchestrator creates a new SagaOrchestrator with the default configuration.
func newSagaOrchestrator() *sagaOrchestrator {
	ctx, cancel := context.WithCancel(context.Background())
	return &sagaOrchestrator{
		mode:                SyncMode,
		steps:               []SagaStep{},
		context:             ctx,
		cancel:              cancel,
		enableExecutionTime: false,
		errCh:               make(chan error),
	}
}

// AddStep adds a step to the Saga.
func (s *sagaOrchestrator) AddStep(action func() (interface{}, error), compensation func(interface{}) error) {
	s.steps = append(s.steps, SagaStep{
		Action:       action,
		Compensation: compensation,
	})
}

// Execute runs the Saga and manages compensations on failure or timeout.
func (s *sagaOrchestrator) Execute() error {
	if s.enableExecutionTime {
		startTime := time.Now()
		defer func() {
			s.ExecutionTime = time.Since(startTime)
			fmt.Printf("Saga execution time: %v\n", s.ExecutionTime)
		}()
	}

	completedSteps := []struct {
		Step     SagaStep
		Response interface{}
	}{}

	if s.mode == AsyncMode {
		var wg sync.WaitGroup

		for _, step := range s.steps {
			wg.Add(1)
			go func(step SagaStep) {
				defer wg.Done()

				select {
				case <-s.context.Done():
					s.errCh <- fmt.Errorf("saga cancelled or timed out")
				default:
					response, err := step.Action()
					if err != nil {
						s.errCh <- err
					} else {
						completedSteps = append(completedSteps, struct {
							Step     SagaStep
							Response interface{}
						}{Step: step, Response: response})
					}
				}
			}(step)
		}

		go func() {
			wg.Wait()
			close(s.errCh)
		}()

		for err := range s.errCh {
			if err != nil {
				// Compensate completed steps
				for i := len(completedSteps) - 1; i >= 0; i-- {
					completedSteps[i].Step.Compensation(completedSteps[i].Response)
				}
				return err
			}
		}
	} else { // SyncMode
		for _, step := range s.steps {
			select {
			case <-s.context.Done():
				// Timeout reached, compensate completed steps
				for i := len(completedSteps) - 1; i >= 0; i-- {
					completedSteps[i].Step.Compensation(completedSteps[i].Response)
				}
				return fmt.Errorf("saga cancelled or timed out")
			default:
				response, err := step.Action()
				if err != nil {
					// Compensate completed steps
					for i := len(completedSteps) - 1; i >= 0; i-- {
						completedSteps[i].Step.Compensation(completedSteps[i].Response)
					}
					return err
				}
				completedSteps = append(completedSteps, struct {
					Step     SagaStep
					Response interface{}
				}{Step: step, Response: response})
			}
		}
	}

	return nil
}

// SetTimeout sets a timeout for the Saga execution.
func (s *sagaOrchestrator) SetTimeout(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	s.context = ctx
	s.cancel = cancel
}

// SetMode sets the execution mode (SyncMode or AsyncMode).
func (s *sagaOrchestrator) SetMode(mode SagaMode) {
	s.mode = mode
}

// EnableExecutionTimeLogging enables or disables execution time logging.
func (s *sagaOrchestrator) EnableExecutionTimeLogging(enable bool) {
	s.enableExecutionTime = enable
}
