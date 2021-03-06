package events

import (
	"github.com/nyaruka/goflow/flows"
)

func init() {
	RegisterType(TypeWebhookCalled, func() flows.Event { return &WebhookCalledEvent{} })
}

// TypeWebhookCalled is the type for our webhook events
const TypeWebhookCalled string = "webhook_called"

// WebhookCalledEvent events are created when a webhook is called. The event contains
// the status and status code of the response, as well as a full dump of the
// request and response. Applying this event updates @run.webhook in the context.
//
//   {
//     "type": "webhook_called",
//     "created_on": "2006-01-02T15:04:05Z",
//     "url": "https://api.ipify.org?format=json",
//     "status": "success",
//     "status_code": 200,
//     "request": "GET https://api.ipify.org?format=json",
//     "response": "HTTP/1.1 200 OK {\"ip\":\"190.154.48.130\"}"
//   }
//
// @event webhook_called
type WebhookCalledEvent struct {
	BaseEvent
	engineOnlyEvent

	URL        string              `json:"url"         validate:"required"`
	Status     flows.WebhookStatus `json:"status"      validate:"required"`
	StatusCode int                 `json:"status_code"`
	Request    string              `json:"request"     validate:"required"`
	Response   string              `json:"response"`
}

// NewWebhookCalledEvent returns a new webhook called event
func NewWebhookCalledEvent(url string, status flows.WebhookStatus, statusCode int, request string, response string) *WebhookCalledEvent {
	return &WebhookCalledEvent{
		BaseEvent:  NewBaseEvent(),
		URL:        url,
		Status:     status,
		StatusCode: statusCode,
		Request:    request,
		Response:   response,
	}
}

// Type returns the type of this event
func (e *WebhookCalledEvent) Type() string { return TypeWebhookCalled }

// Apply applies this event to the given run
func (e *WebhookCalledEvent) Apply(run flows.FlowRun) error {
	run.SetWebhook(flows.NewWebhookCall(e.URL, e.Status, e.StatusCode, e.Request, e.Response))
	return nil
}
