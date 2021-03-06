package actions

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nyaruka/goflow/flows"
	"github.com/nyaruka/goflow/flows/events"
)

func init() {
	RegisterType(TypeCallWebhook, func() flows.Action { return &CallWebhookAction{} })
}

// TypeCallWebhook is the type for the call webhook action
const TypeCallWebhook string = "call_webhook"

// CallWebhookAction can be used to call an external service. The body, header and url fields may be
// templates and will be evaluated at runtime.
//
// A [event:webhook_called] event will be created based on the results of the HTTP call.
//
//   {
//     "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
//     "type": "call_webhook",
//     "method": "GET",
//     "url": "http://localhost:49998/?cmd=success",
//     "headers": {
//       "Authorization": "Token AAFFZZHH"
//     }
//   }
//
// @action call_webhook
type CallWebhookAction struct {
	BaseAction
	onlineAction

	Method  string            `json:"method"             validate:"required,http_method"`
	URL     string            `json:"url"                validate:"required"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
}

// Type returns the type of this action
func (a *CallWebhookAction) Type() string { return TypeCallWebhook }

// Validate validates our action is valid and has all the assets it needs
func (a *CallWebhookAction) Validate(assets flows.SessionAssets) error {
	if a.Body != "" && a.Method == "GET" {
		return fmt.Errorf("can specifiy body if method is GET")
	}

	return nil
}

// Execute runs this action
func (a *CallWebhookAction) Execute(run flows.FlowRun, step flows.Step, log flows.EventLog) error {

	// substitute any variables in our url
	url, err := run.EvaluateTemplateAsString(a.URL, true)
	if err != nil {
		log.Add(events.NewErrorEvent(err))
	}
	if url == "" {
		log.Add(events.NewErrorEvent(fmt.Errorf("call_webhook URL evaluated to empty string, skipping")))
		return nil
	}

	method := strings.ToUpper(a.Method)
	body := a.Body

	// substitute any body variables
	if body != "" {
		body, err = run.EvaluateTemplateAsString(body, false)
		if err != nil {
			log.Add(events.NewErrorEvent(err))
		}
	}

	// build our request
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		log.Add(events.NewErrorEvent(err))
		return nil
	}

	// add the custom headers, substituting any template vars
	for key, value := range a.Headers {
		headerValue, err := run.EvaluateTemplateAsString(value, false)
		if err != nil {
			log.Add(events.NewErrorEvent(err))
		}

		req.Header.Add(key, headerValue)
	}

	webhook, err := flows.MakeWebhookCall(run.Session(), req)

	if err != nil {
		log.Add(events.NewErrorEvent(err))
	} else {
		log.Add(events.NewWebhookCalledEvent(webhook.URL(), webhook.Status(), webhook.StatusCode(), webhook.Request(), webhook.Response()))
	}

	return nil
}
