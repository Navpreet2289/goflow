package events

import (
	"fmt"

	"github.com/nyaruka/goflow/flows"
)

func init() {
	RegisterType(TypeRunExpired, func() flows.Event { return &RunExpiredEvent{} })
}

// TypeRunExpired is the type of our flow expired event
const TypeRunExpired string = "run_expired"

// RunExpiredEvent events are sent by the caller to notify the engine that a run has expired
//
//   {
//     "type": "run_expired",
//     "created_on": "2006-01-02T15:04:05Z",
//     "run_uuid": "0e06f977-cbb7-475f-9d0b-a0c4aaec7f6a"
//   }
//
// @event run_expired
type RunExpiredEvent struct {
	BaseEvent
	callerOnlyEvent

	RunUUID flows.RunUUID `json:"run_uuid"    validate:"required,uuid4"`
}

// Type returns the type of this event
func (e *RunExpiredEvent) Type() string { return TypeRunExpired }

// Validate validates our event is valid and has all the assets it needs
func (e *RunExpiredEvent) Validate(assets flows.SessionAssets) error {
	return nil
}

// Apply applies this event to the given run
func (e *RunExpiredEvent) Apply(run flows.FlowRun) error {
	if run.UUID() != e.RunUUID {
		return fmt.Errorf("only the current run can be expired")
	}

	run.Exit(flows.RunStatusExpired)
	return nil
}
