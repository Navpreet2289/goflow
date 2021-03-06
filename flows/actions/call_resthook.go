package actions

import (
	"net/http"
	"strings"

	"github.com/nyaruka/goflow/assets/rest/types"
	"github.com/nyaruka/goflow/flows"
	"github.com/nyaruka/goflow/flows/events"
)

func init() {
	RegisterType(TypeCallResthook, func() flows.Action { return &CallResthookAction{} })
}

// TypeCallResthook is the type for the call resthook action
const TypeCallResthook string = "call_resthook"

// CallResthookAction can be used to call a resthook.
//
// A [event:resthook_called] event will be created based on the results of the HTTP call
// to each subscriber of the resthook.
//
//   {
//     "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
//     "type": "call_resthook",
//     "resthook": "new-registration"
//   }
//
// @action call_resthook
type CallResthookAction struct {
	BaseAction
	onlineAction

	Resthook string `json:"resthook" validate:"required"`
}

// Type returns the type of this action
func (a *CallResthookAction) Type() string { return TypeCallResthook }

// Validate validates our action is valid and has all the assets it needs
func (a *CallResthookAction) Validate(assets flows.SessionAssets) error {
	return nil
}

// Execute runs this action
func (a *CallResthookAction) Execute(run flows.FlowRun, step flows.Step, log flows.EventLog) error {
	// if resthook doesn't exist, treat it like an existing one with no subscribers
	resthook := run.Session().Assets().Resthooks().FindBySlug(a.Resthook)
	if resthook == nil {
		resthook = flows.NewResthook(types.NewResthook(a.Resthook, nil))
	}

	// build our payload
	payload, err := run.EvaluateTemplateAsString(flows.DefaultWebhookPayload, false)
	if err != nil {
		log.Add(events.NewErrorEvent(err))
	}

	// make a call to each subscriber URL
	calls := make([]*events.ResthookSubscriberCall, 0, len(resthook.Subscribers()))

	for _, url := range resthook.Subscribers() {
		req, err := http.NewRequest("POST", url, strings.NewReader(payload))
		if err != nil {
			log.Add(events.NewErrorEvent(err))
			return nil
		}

		req.Header.Add("Content-Type", "application/json")

		webhook, err := flows.MakeWebhookCall(run.Session(), req)
		if err != nil {
			log.Add(events.NewErrorEvent(err))
		} else {
			calls = append(calls, events.NewResthookSubscriberCall(webhook))
		}
	}

	log.Add(events.NewResthookCalledEvent(a.Resthook, payload, calls))

	return nil
}
