package driver

import (
	"fmt"
	"testing"
	"time"

	uerr "github.com/uspeedo/uspeedo-sdk-go/uspeedo/error"
)

type StepReport struct {
	Title     string          `json:"title"`
	Status    string          `json:"status"`
	Execution StepExecution   `json:"execution"`
	Retries   Retries         `json:"retries"`
	Errors    executionErrors `json:"errors,omitempty"`
}

type StepExecution struct {
	MaxRetries    int     `json:"max_retries"`
	RetryInterval float64 `json:"retry_interval"`
	StartupDelay  float64 `json:"startup_delay"`
	FastFail      bool    `json:"fast_fail"`
	Duration      float64 `json:"duration"`
	StartTime     float64 `json:"start_time"`
	EndTime       float64 `json:"end_time"`
}

type TestValidator func(interface{}) error
type Step struct {
	T *testing.T

	MaxRetries    int
	RetryInterval time.Duration
	StartupDelay  time.Duration
	FastFail      bool
	Title         string
	Scenario      *Scenario
	retries       Retries

	Invoker    func(*Step) (interface{}, error)
	Validators func(*Step) []TestValidator

	errors    []error
	startTime float64
	endTime   float64
	status    string
	id        int
}

type Retries struct {
	Headers []string        `json:"headers"`
	Rows    [][]interface{} `json:"rows"`
}

func (step *Step) AppendRetriesRows(row []interface{}) {
	step.retries.Rows = append(step.retries.Rows, row)
}

func (step *Step) SetRetriesHeaders(headers []string) {
	step.retries.Headers = headers
}

// LoadFixture is a function for load fixture by the name from map fixture function
func (step *Step) LoadFixture(name string) (interface{}, error) {
	if step.Scenario.Spec.fixtures[name] != nil {
		return step.Scenario.Spec.fixtures[name](step)
	}
	return nil, fmt.Errorf("can not load fixture by the %s", name)
}

// Must will check error is nil and return the value
func (step *Step) Must(v interface{}, err error) interface{} {
	if err != nil {
		step.AppendError(err)
	}
	return v
}

func (step *Step) AppendError(err error) {
	step.errors = append(step.errors, fmt.Errorf("step %02d Failed, %s", step.id, err))
}

// Run will run the step test case with retry
func (step *Step) run() {
	step.startTime = float64(time.Now().Unix())
	if step.StartupDelay != time.Duration(0) {
		<-time.After(step.StartupDelay)
	}

	defer func() {
		step.endTime = float64(time.Now().Unix())
	}()

	for i := 0; i < step.MaxRetries+1; i++ {
		step.errors = []error{}

		resp, err := step.Invoker(step)
		if err != nil {
			if e, ok := err.(uerr.Error); ok && e.Name() == uerr.ErrSendRequest {
				step.status = "failed"
				step.AppendError(err)
				return
			} else if ok && e.Name() == uerr.ErrRetCode {
				// pass
			} else {
				step.AppendError(err)
				// continue
			}
		}

		if step.Validators != nil && resp != nil {
			for _, validator := range step.Validators(step) {
				if err := validator(resp); err != nil {
					step.AppendError(err)
				}
			}
		}

		if len(step.errors) > 0 {
			if i == step.MaxRetries {
				step.status = "failed"
				return
			}
			<-time.After(step.RetryInterval)
			continue
		}

		step.status = "passed"
		return
	}
}

func (step *Step) Report() StepReport {
	return StepReport{
		Title:  step.Title,
		Status: step.status,
		Execution: StepExecution{
			MaxRetries:    step.MaxRetries,
			RetryInterval: step.RetryInterval.Seconds(),
			StartupDelay:  step.StartupDelay.Seconds(),
			FastFail:      step.FastFail,
			Duration:      step.endTime - step.startTime,
			StartTime:     step.startTime,
			EndTime:       step.endTime,
		},
		Retries: step.retries,
		Errors:  step.errors,
	}
}

func (step *Step) init() {
	step.status = "skipped"
}
