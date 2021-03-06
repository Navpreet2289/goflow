package events

import (
	"fmt"

	"github.com/nyaruka/goflow/assets"
	"github.com/nyaruka/goflow/flows"
)

func init() {
	RegisterType(TypeContactGroupsRemoved, func() flows.Event { return &ContactGroupsRemovedEvent{} })
}

// TypeContactGroupsRemoved is the type fo our remove from group action
const TypeContactGroupsRemoved string = "contact_groups_removed"

// ContactGroupsRemovedEvent events are created when a contact has been removed from one or more
// groups.
//
//   {
//     "type": "contact_groups_removed",
//     "created_on": "2006-01-02T15:04:05Z",
//     "groups": [{"uuid": "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d", "name": "Reporters"}]
//   }
//
// @event contact_groups_removed
type ContactGroupsRemovedEvent struct {
	BaseEvent
	callerOrEngineEvent

	Groups []*assets.GroupReference `json:"groups" validate:"required,min=1,dive"`
}

// NewContactGroupsRemovedEvent returns a new remove from group event
func NewContactGroupsRemovedEvent(groups []*assets.GroupReference) *ContactGroupsRemovedEvent {
	return &ContactGroupsRemovedEvent{
		BaseEvent: NewBaseEvent(),
		Groups:    groups,
	}
}

// Type returns the type of this event
func (e *ContactGroupsRemovedEvent) Type() string { return TypeContactGroupsRemoved }

// Validate validates our event is valid and has all the assets it needs
func (e *ContactGroupsRemovedEvent) Validate(assets flows.SessionAssets) error {
	for _, group := range e.Groups {
		if _, err := assets.Groups().Get(group.UUID); err != nil {
			return err
		}
	}
	return nil
}

// Apply applies this event to the given run
func (e *ContactGroupsRemovedEvent) Apply(run flows.FlowRun) error {
	if run.Contact() == nil {
		return fmt.Errorf("can't apply event in session without a contact")
	}

	assets := run.Session().Assets()

	for _, groupRef := range e.Groups {
		group, err := assets.Groups().Get(groupRef.UUID)
		if err == nil {
			run.Contact().Groups().Remove(group)
		}
	}
	return nil
}
