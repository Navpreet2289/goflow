package events

import (
	"fmt"
	"time"

	"github.com/nyaruka/goflow/flows"
	"github.com/nyaruka/goflow/utils"
)

var registeredTypes = map[string](func() flows.Event){}

// RegisterType registers a new type of router
func RegisterType(name string, initFunc func() flows.Event) {
	registeredTypes[name] = initFunc
}

// BaseEvent is the base of all event types
type BaseEvent struct {
	CreatedOn_  time.Time      `json:"created_on" validate:"required"`
	StepUUID_   flows.StepUUID `json:"step_uuid,omitempty" validate:"omitempty,uuid4"`
	FromCaller_ bool           `json:"-"`
}

// NewBaseEvent creates a new base event
func NewBaseEvent() BaseEvent {
	return BaseEvent{CreatedOn_: utils.Now()}
}

// CreatedOn returns the created on time of this event
func (e *BaseEvent) CreatedOn() time.Time { return e.CreatedOn_ }

// StepUUID returns the UUID of the step in the path where this event occured
func (e *BaseEvent) StepUUID() flows.StepUUID { return e.StepUUID_ }

// SetStepUUID sets the UUID of the step in the path where this event occured
func (e *BaseEvent) SetStepUUID(stepUUID flows.StepUUID) { e.StepUUID_ = stepUUID }

// FromCaller returns whether this event originated from the caller
func (e *BaseEvent) FromCaller() bool { return e.FromCaller_ }

// SetFromCaller sets whether this event originated from the caller
func (e *BaseEvent) SetFromCaller(fromCaller bool) { e.FromCaller_ = fromCaller }

// utility struct which sets the allowed origin of an event to "caller-only"
type callerOnlyEvent struct{}

// AllowedOrigin determines where this event type can originate
func (e *callerOnlyEvent) AllowedOrigin() flows.EventOrigin { return flows.EventOriginCaller }

// utility struct which sets the allowed origin of an event to "engine-only"
type engineOnlyEvent struct{}

// AllowedOrigin determines where this event type can originate
func (e *engineOnlyEvent) AllowedOrigin() flows.EventOrigin { return flows.EventOriginEngine }

// Validate validates our event is valid and has all the assets it needs. We assume engine generated events are valid.
func (e *engineOnlyEvent) Validate(assets flows.SessionAssets) error {
	return nil
}

// utility struct which sets the allowed origin of an event to "engine or caller"
type callerOrEngineEvent struct{}

// AllowedOrigin determines where this event type can originate
func (e *callerOrEngineEvent) AllowedOrigin() flows.EventOrigin {
	return flows.EventOriginCaller | flows.EventOriginEngine
}

//------------------------------------------------------------------------------------------
// JSON Encoding / Decoding
//------------------------------------------------------------------------------------------

// ReadEvent reads a single event from the given envelope
func ReadEvent(envelope *utils.TypedEnvelope) (flows.Event, error) {
	f := registeredTypes[envelope.Type]
	if f == nil {
		return nil, fmt.Errorf("unknown type: %s", envelope.Type)
	}

	event := f()
	return event, utils.UnmarshalAndValidate(envelope.Data, event)
}

// ReadEvents reads the events from the given envelopes
func ReadEvents(envelopes []*utils.TypedEnvelope) ([]flows.Event, error) {
	events := make([]flows.Event, len(envelopes))
	for e, envelope := range envelopes {
		event, err := ReadEvent(envelope)
		if err != nil {
			return nil, fmt.Errorf("unable to read event[type=%s]: %s", envelope.Type, err)
		}
		event.SetFromCaller(true)
		events[e] = event
	}
	return events, nil
}

// EventsToEnvelopes converts the given events to typed envelopes
func EventsToEnvelopes(events []flows.Event) ([]*utils.TypedEnvelope, error) {
	var err error
	envelopes := make([]*utils.TypedEnvelope, len(events))
	for e, event := range events {
		if envelopes[e], err = utils.EnvelopeFromTyped(event); err != nil {
			return nil, err
		}
	}
	return envelopes, nil
}
