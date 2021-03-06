# Container

Flow definitions are defined as a list of nodes, the first node being the entry into the flow. The simplest possible flow containing no nodes whatsoever (and therefore being a no-op) contains the following fields:

 * `uuid` the UUID
 * `name` the name
 * `language` the base authoring language used for localization
 * `type` the type, one of `messaging`, `messaging_offline`, `voice` whcih determines which actions are allowed in the flow
 * `nodes` the nodes (may be empty)

For example:

```json
{
    "uuid": "b7bb5e7c-ad49-4e65-9e24-bf7f1e4ff00a",
    "name": "Empty Flow",
    "language": "eng",
    "type": "messaging",
    "nodes": []
}
```

# Nodes

Flow definitions are composed of zero or more nodes, the first node is always the entry node.

A Node consists of:

 * `uuid` the UUID
 * `actions` a list of 0-n actions which will be executed upon first entering a node
 * `wait` an optional pause in the flow waiting for some event to occur, such as a contact responding, a timeout for that response or a subflow completing
 * `exit` a list of 0-n exits which can be used to link to other nodes
 * `router` an optional router which determines which exit to take

At its simplest, a node can be just a single action with no exits, wait or router, such as:

```json
{
    "uuid":"5a06445e-d790-4bd3-a10b-b47bdcc9abed",
    "actions":[{
        "uuid": "abc0a2bf-6b4a-4ee0-83e1-1eebae6948ac",
        "type": "send_msg",
        "text": "What is your name?"
    }]
}
```

If a node wishes to route to another node, it can do so by defining one or more exits, each with the UUID of the node that is next. Without a router defined, the first exit will always be taken. 

An exit consists of:

 * `uuid` the UUID
 * `destination_node_uuid` the uuid of the node that should be visited if this exit is chosen by the router (optional)
 * `name` a name for this exit (optional)

```json
{
    "uuid":"5a06445e-d790-4bd3-a10b-b47bdcc9abed",
    "actions":[{
        "uuid": "abc0a2bf-6b4a-4ee0-83e1-1eebae6948ac",
        "type": "send_msg",
        "text": "What is your name?"
    }],
    "exits": [{
        "uuid":"eb7defc9-3c66-4dfc-80bc-825567ccd9de",
        "destination_node_uuid":"ee0bee3f-34b3-4275-af78-f9ff52c82e6a"
    }]
}
```

# Routers

## Switch

If a node wishes to route differently based on some state, it can add a `switch` router which defines one or more `cases`. 
Each case defines a `type` which is the name of an expression function that is run by passing the evaluation of `operand` 
as the first argument. Cases may define additional arguments using the `arguments` array on a case. If no case evaluates 
to true, then the `default_exit_uuid` will be used otherwise flow execution will stop.

A switch router may also define a `result_name` parameters which will save the result of the case which evaluated as true.

A switch router consists of:

 * `operand` the expression which will be evaluated against each of our cases
 * `default_exit_uuid` the uuid of the default exit to take if no case matches (optional)
 * `result_name` the name of the result which should be written when the switch is evaluated (optional)
 * `cases` a list of 1-n cases which are evaluated in order until one is true

Each case consists of:

 * `uuid` the UUID
 * `type` the type of this test, this must be an excellent test (see below) and will be passed the value of the switch's operand as its first value
 * `arguments` an optional list of templates which can be passed as extra parameters to the test (after the initial operand)
 * `exit_uuid` the uuid of the exit that should be taken if this case evaluated to true

 An example switch router that tests for the input not being empty:

```json
{
    "uuid":"ee0bee3f-34b3-4275-af78-f9ff52c82e6a",
    "router": {
        "type":"switch",
        "operand": "@run.input",
        "default_exit_uuid": "9574fbfd-510f-4dfc-b989-97d2aecf50b9",
        "cases": [{
            "uuid": "6f78d564-029b-4715-b8d4-b28daeae4f24",
            "type": "has_text",
            "exit_uuid": "cab600f5-b54b-49b9-a7ea-5638f4cbf2b4"
        }]
    },
    "exits": [{
        "uuid":"cab600f5-b54b-49b9-a7ea-5638f4cbf2b4",
        "name":"Has Name",
        "destination_node_uuid":"deec1dd4-b727-4b21-800a-0b7bbd146a82"
    },{
        "uuid":"9574fbfd-510f-4dfc-b989-97d2aecf50b9",
        "name":"Other",
        "destination_node_uuid":"ee0bee3f-34b3-4275-af78-f9ff52c82e6a"
    }]
}
```

# Waits

A node can indicate that it needs more information to continue by containing a wait.

## Msg

This wait type indicates that flow execution should pause until an incoming message is received and also gives an optional timeout in seconds as to when the flow 
should continue even if there is no reply:

```json
{
    "type": "msg",
    "timeout": 600
}
```

## Nothing

This wait type indicates that the caller can resume the session immediately with no incoming message or any other input. This type of
wait enables the caller to commit changes in the session up to that point in the flow.

```json
{
    "type": "nothing"
}
```

# Actions

Actions on a node generate events which can then be ingested by the engine container. In some cases the actions cause an immediate action, such 
as calling a webhook, in others the engine container is responsible for taking the action based on the event that is output, such as sending 
messages or updating contact fields. In either case the internal state of the engine is always updated to represent the new state so that
flow execution is consistent. For example, while the engine itself does not have access to a contact store, it updates its internal 
representation of a contact's state based on action performed on a flow so that later references in the flow are correct.

<div class="actions">
<a name="action:add_contact_groups"></a>

## add_contact_groups

Can be used to add a contact to one or more groups. A [contact_groups_added](sessions.html#event:contact_groups_added) event will be created
for the groups which the contact has been added to.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "add_contact_groups",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "groups": [
        {
            "uuid": "1e1ce1e1-9288-4504-869e-022d1003c72a",
            "name": "Customers"
        }
    ]
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "contact_groups_added",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "4f15f627-b1e2-4851-8dbf-00ecf5d03034",
    "groups": [
        {
            "uuid": "1e1ce1e1-9288-4504-869e-022d1003c72a",
            "name": "Customers"
        }
    ]
}
```
</div>
<a name="action:add_contact_urn"></a>

## add_contact_urn

Can be used to add a URN to the current contact. A [contact_urn_added](sessions.html#event:contact_urn_added) event
will be created when this action is encountered. If there is no contact then this
action will be ignored.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "add_contact_urn",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "scheme": "tel",
    "path": "@run.results.phone_number"
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "contact_urn_added",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "b504fe9e-d8a8-47fd-af9c-ff2f1faac4db",
    "urn": "tel:+12344563452"
}
```
</div>
<a name="action:add_input_labels"></a>

## add_input_labels

Can be used to add labels to the last user input on a flow. An [input_labels_added](sessions.html#event:input_labels_added) event
will be created with the labels added when this action is encountered. If there is
no user input at that point then this action will be ignored.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "add_input_labels",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "labels": [
        {
            "uuid": "3f65d88a-95dc-4140-9451-943e94e06fea",
            "name": "Spam"
        }
    ]
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "input_labels_added",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "f3cbd795-9bb3-4331-ba82-c15b24dd577f",
    "input_uuid": "9bf91c2b-ce58-4cef-aacc-281e03f69ab5",
    "labels": [
        {
            "uuid": "3f65d88a-95dc-4140-9451-943e94e06fea",
            "name": "Spam"
        }
    ]
}
```
</div>
<a name="action:call_resthook"></a>

## call_resthook

Can be used to call a resthook.

A [resthook_called](sessions.html#event:resthook_called) event will be created based on the results of the HTTP call
to each subscriber of the resthook.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "call_resthook",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "resthook": "new-registration"
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "resthook_called",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "229bd432-dac7-4a3f-ba91-c48ad8c50e6b",
    "resthook": "new-registration",
    "payload": "{\n\t\"contact\": {\"uuid\": \"5d76d86b-3bb9-4d5a-b822-c9d86f5d8e4f\", \"name\": \"Ryan Lewis\", \"urn\": \"tel:+12065551212\"},\n\t\"flow\": {\"name\":\"Registration\",\"revision\":123,\"uuid\":\"50c3706e-fedb-42c0-8eab-dda3335714b7\"},\n\t\"path\": [{\"arrived_on\":\"2018-04-11T18:24:30.123456Z\",\"exit_uuid\":\"37d8813f-1402-4ad2-9cc2-e9054a96525b\",\"node_uuid\":\"72a1f5df-49f9-45df-94c9-d86f7ea064e5\",\"uuid\":\"347b55be-7be1-4e68-aaa3-04d3fbce5f9a\"},{\"arrived_on\":\"2018-04-11T18:24:30.123456Z\",\"exit_uuid\":\"d898f9a4-f0fc-4ac4-a639-c98c602bb511\",\"node_uuid\":\"f5bb9b7a-7b5e-45c3-8f0e-61b4e95edf03\",\"uuid\":\"da339edd-083b-48cb-bef6-3979f99a96f9\"},{\"arrived_on\":\"2018-04-11T18:24:30.123456Z\",\"exit_uuid\":\"\",\"node_uuid\":\"c0781400-737f-4940-9a6c-1ec1c3df0325\",\"uuid\":\"229bd432-dac7-4a3f-ba91-c48ad8c50e6b\"}],\n\t\"results\": {\"favorite_color\":{\"category\":\"Red\",\"category_localized\":\"Red\",\"created_on\":\"2018-04-11T18:24:30.123456Z\",\"input\":null,\"name\":\"Favorite Color\",\"node_uuid\":\"f5bb9b7a-7b5e-45c3-8f0e-61b4e95edf03\",\"value\":\"red\"},\"phone_number\":{\"category\":\"\",\"category_localized\":\"\",\"created_on\":\"2018-04-11T18:24:30.123456Z\",\"input\":null,\"name\":\"Phone Number\",\"node_uuid\":\"f5bb9b7a-7b5e-45c3-8f0e-61b4e95edf03\",\"value\":\"+12344563452\"}},\n\t\"run\": {\"uuid\": \"4c9abf31-d821-4e97-ba7e-53c2263e32f8\", \"created_on\": \"2018-04-11T18:24:30.123456Z\"},\n\t\"input\": {\"attachments\":[{\"content_type\":\"image/jpeg\",\"url\":\"http://s3.amazon.com/bucket/test.jpg\"},{\"content_type\":\"audio/mp3\",\"url\":\"http://s3.amazon.com/bucket/test.mp3\"}],\"channel\":{\"address\":\"+12345671111\",\"name\":\"My Android Phone\",\"uuid\":\"57f1078f-88aa-46f4-a59a-948a5739c03d\"},\"created_on\":\"2000-01-01T00:00:00.000000Z\",\"text\":\"Hi there\",\"type\":\"msg\",\"urn\":{\"display\":\"\",\"path\":\"+12065551212\",\"scheme\":\"tel\"},\"uuid\":\"9bf91c2b-ce58-4cef-aacc-281e03f69ab5\"},\n\t\"channel\": {\"address\":\"+12345671111\",\"name\":\"My Android Phone\",\"uuid\":\"57f1078f-88aa-46f4-a59a-948a5739c03d\"}\n}",
    "calls": [
        {
            "url": "http://127.0.0.1:49998/?cmd=success",
            "status": "success",
            "status_code": 200,
            "response": "HTTP/1.1 200 OK\r\nContent-Length: 16\r\nContent-Type: text/plain; charset=utf-8\r\nDate: Wed, 11 Apr 2018 18:24:30 GMT\r\n\r\n{ \"ok\": \"true\" }"
        }
    ]
}
```
</div>
<a name="action:call_webhook"></a>

## call_webhook

Can be used to call an external service. The body, header and url fields may be
templates and will be evaluated at runtime.

A [webhook_called](sessions.html#event:webhook_called) event will be created based on the results of the HTTP call.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "call_webhook",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "method": "GET",
    "url": "http://localhost:49998/?cmd=success",
    "headers": {
        "Authorization": "Token AAFFZZHH"
    }
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "webhook_called",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "e68a851e-6328-426b-a8fd-1537ca860f97",
    "url": "http://localhost:49998/?cmd=success",
    "status": "success",
    "status_code": 200,
    "request": "GET /?cmd=success HTTP/1.1\r\nHost: localhost:49998\r\nUser-Agent: goflow-testing\r\nAuthorization: Token AAFFZZHH\r\nAccept-Encoding: gzip\r\n\r\n",
    "response": "HTTP/1.1 200 OK\r\nContent-Length: 16\r\nContent-Type: text/plain; charset=utf-8\r\nDate: Wed, 11 Apr 2018 18:24:30 GMT\r\n\r\n{ \"ok\": \"true\" }"
}
```
</div>
<a name="action:remove_contact_groups"></a>

## remove_contact_groups

Can be used to remove a contact from one or more groups. A [contact_groups_removed](sessions.html#event:contact_groups_removed) event will be created
for the groups which the contact is removed from. Groups can either be explicitly provided or `all_groups` can be set to true to remove
the contact from all non-dynamic groups.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "remove_contact_groups",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "groups": [
        {
            "uuid": "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d",
            "name": "Registered Users"
        }
    ],
    "all_groups": false
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "contact_groups_removed",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "5fa51f39-76ea-421c-a71b-fe4af29b871a",
    "groups": [
        {
            "uuid": "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d",
            "name": "Testers"
        }
    ]
}
```
</div>
<a name="action:send_broadcast"></a>

## send_broadcast

Can be used to send a message to one or more contacts. It accepts a list of URNs, a list of groups
and a list of contacts.

The URNs and text fields may be templates. A [broadcast_created](sessions.html#event:broadcast_created) event will be created for each unique urn, contact and group
with the evaluated text.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "send_broadcast",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "text": "Hi @contact.name, are you ready to complete today's survey?",
    "attachments": null,
    "urns": [
        "tel:+12065551212"
    ]
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "broadcast_created",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "8e64b588-d46e-4016-a5ef-59cf4d9d7a5b",
    "translations": {
        "eng": {
            "text": "Hi Ryan Lewis, are you ready to complete today's survey?"
        }
    },
    "base_language": "eng",
    "urns": [
        "tel:+12065551212"
    ]
}
```
</div>
<a name="action:send_email"></a>

## send_email

Can be used to send an email to one or more recipients. The subject, body and addresses
can all contain expressions.

An [email_created](sessions.html#event:email_created) event will be created for each email address.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "send_email",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "addresses": [
        "@contact.urns.mailto.0"
    ],
    "subject": "Here is your activation token",
    "body": "Your activation token is @contact.fields.activation_token"
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "email_created",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "08eba586-0bb1-47ab-8c15-15a7c0c5228d",
    "addresses": [
        "foo@bar.com"
    ],
    "subject": "Here is your activation token",
    "body": "Your activation token is AACC55"
}
```
</div>
<a name="action:send_msg"></a>

## send_msg

Can be used to reply to the current contact in a flow. The text field may contain templates.

A [msg_created](sessions.html#event:msg_created) event will be created with the evaluated text.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "send_msg",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "text": "Hi @contact.name, are you ready to complete today's survey?",
    "attachments": []
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "msg_created",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "c1f115c7-bcf3-44ef-88b2-5d345629f07f",
    "msg": {
        "uuid": "10c62052-7db1-49d1-b8ba-60d66db82e39",
        "urn": "tel:+12065551212?channel=57f1078f-88aa-46f4-a59a-948a5739c03d",
        "channel": {
            "uuid": "57f1078f-88aa-46f4-a59a-948a5739c03d",
            "name": "My Android Phone"
        },
        "text": "Hi Ryan Lewis, are you ready to complete today's survey?"
    }
}
```
</div>
<a name="action:set_contact_channel"></a>

## set_contact_channel

Can be used to update the preferred channel of the current contact.

A [contact_channel_changed](sessions.html#event:contact_channel_changed) event will be created with the set channel.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "set_contact_channel",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "channel": {
        "uuid": "4bb288a0-7fca-4da1-abe8-59a593aff648",
        "name": "FAcebook Channel"
    }
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "contact_channel_changed",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "c174a241-6057-41a3-874b-f17fb8365c22",
    "channel": {
        "uuid": "4bb288a0-7fca-4da1-abe8-59a593aff648",
        "name": "FAcebook Channel"
    }
}
```
</div>
<a name="action:set_contact_field"></a>

## set_contact_field

Can be used to update a field value on the contact. The value is a localizable
template and white space is trimmed from the final value. An empty string clears the value.
A [contact_field_changed](sessions.html#event:contact_field_changed) event will be created with the corresponding value.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "set_contact_field",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "field": {
        "key": "gender",
        "name": "Gender"
    },
    "value": "Male"
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "contact_field_changed",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "a08b46fc-f057-4e9a-9bd7-277a6a165264",
    "field": {
        "key": "gender",
        "name": "Gender"
    },
    "value": "Male"
}
```
</div>
<a name="action:set_contact_language"></a>

## set_contact_language

Can be used to update the name of the contact. The language is a localizable
template and white space is trimmed from the final value. An empty string clears the language.
A [contact_language_changed](sessions.html#event:contact_language_changed) event will be created with the corresponding value.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "set_contact_language",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "language": "eng"
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "contact_language_changed",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "7ca3fc1e-e652-4f5c-979e-17606f578787",
    "language": "eng"
}
```
</div>
<a name="action:set_contact_name"></a>

## set_contact_name

Can be used to update the name of the contact. The name is a localizable
template and white space is trimmed from the final value. An empty string clears the name.
A [contact_name_changed](sessions.html#event:contact_name_changed) event will be created with the corresponding value.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "set_contact_name",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "name": "Bob Smith"
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "contact_name_changed",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "fbce9f1c-ddff-45f4-8d46-86b76f70a6a6",
    "name": "Bob Smith"
}
```
</div>
<a name="action:set_contact_timezone"></a>

## set_contact_timezone

Can be used to update the timezone of the contact. The timezone is a localizable
template and white space is trimmed from the final value. An empty string clears the timezone.
A [contact_timezone_changed](sessions.html#event:contact_timezone_changed) event will be created with the corresponding value.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "set_contact_timezone",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "timezone": "Africa/Kigali"
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "contact_timezone_changed",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "e4be9d25-b3ab-4a47-8704-ab259cb52a5d",
    "timezone": "Africa/Kigali"
}
```
</div>
<a name="action:set_run_result"></a>

## set_run_result

Can be used to save a result for a flow. The result will be available in the context
for the run as @run.results.[name]. The optional category can be used as a way of categorizing results,
this can be useful for reporting or analytics.

Both the value and category fields may be templates. A [run_result_changed](sessions.html#event:run_result_changed) event will be created with the
final values.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "set_run_result",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "name": "Gender",
    "value": "m",
    "category": "Male"
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "run_result_changed",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "bb7de8fc-d0b0-41a6-bdf0-950b64bbbc6d",
    "name": "Gender",
    "value": "m",
    "category": "Male",
    "node_uuid": "c0781400-737f-4940-9a6c-1ec1c3df0325"
}
```
</div>
<a name="action:start_flow"></a>

## start_flow

Can be used to start a contact down another flow. The current flow will pause until the subflow exits or expires.

A [flow_triggered](sessions.html#event:flow_triggered) event will be created to record that the flow was started.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "start_flow",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "flow": {
        "uuid": "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d",
        "name": "Collect Language"
    }
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "flow_triggered",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "dda50da0-8fc0-4f22-9c96-61ebc05df996",
    "flow": {
        "uuid": "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d",
        "name": "Collect Language"
    },
    "parent_run_uuid": "a8ff08ef-6f27-44bd-9029-066bfcb36cf8"
}
```
</div>
<a name="action:start_session"></a>

## start_session

Can be used to trigger sessions for other contacts and groups. A [session_triggered](sessions.html#event:session_triggered) event
will be created and it's the responsibility of the caller to act on that by initiating a new session with the flow engine.

<div class="input_action"><h3>Action</h3>```json
{
    "type": "start_session",
    "uuid": "8eebd020-1af5-431c-b943-aa670fc74da9",
    "groups": [
        {
            "uuid": "1e1ce1e1-9288-4504-869e-022d1003c72a",
            "name": "Customers"
        }
    ],
    "flow": {
        "uuid": "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d",
        "name": "Registration"
    }
}
```
</div><div class="output_event"><h3>Event</h3>```json
{
    "type": "session_triggered",
    "created_on": "2018-04-11T18:24:30.123456Z",
    "step_uuid": "636bcfe8-1dd9-4bbd-a2a5-6b6ffeeada26",
    "flow": {
        "uuid": "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d",
        "name": "Registration"
    },
    "groups": [
        {
            "uuid": "1e1ce1e1-9288-4504-869e-022d1003c72a",
            "name": "Customers"
        }
    ],
    "run": {
        "uuid": "e6e30b78-f9c1-462b-9418-6d3e4ae5a100",
        "flow": {
            "uuid": "50c3706e-fedb-42c0-8eab-dda3335714b7",
            "name": "Registration"
        },
        "contact": {
            "uuid": "5d76d86b-3bb9-4d5a-b822-c9d86f5d8e4f",
            "id": 1234567,
            "name": "Ryan Lewis",
            "language": "eng",
            "timezone": "America/Guayaquil",
            "created_on": "2018-06-20T11:40:30.123456789Z",
            "urns": [
                "tel:+12065551212?channel=57f1078f-88aa-46f4-a59a-948a5739c03d",
                "twitterid:54784326227#nyaruka",
                "mailto:foo@bar.com"
            ],
            "groups": [
                {
                    "uuid": "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d",
                    "name": "Testers"
                },
                {
                    "uuid": "4f1f98fc-27a7-4a69-bbdb-24744ba739a9",
                    "name": "Males"
                }
            ],
            "fields": {
                "activation_token": {
                    "text": "AACC55"
                },
                "age": {
                    "text": "23",
                    "number": 23
                },
                "gender": {
                    "text": "Male"
                },
                "join_date": {
                    "text": "2017-12-02",
                    "datetime": "2017-12-02T00:00:00-02:00"
                }
            }
        },
        "status": "active",
        "results": {
            "favorite_color": {
                "name": "Favorite Color",
                "value": "red",
                "category": "Red",
                "node_uuid": "f5bb9b7a-7b5e-45c3-8f0e-61b4e95edf03",
                "created_on": "2018-04-11T18:24:30.123456Z"
            },
            "phone_number": {
                "name": "Phone Number",
                "value": "+12344563452",
                "node_uuid": "f5bb9b7a-7b5e-45c3-8f0e-61b4e95edf03",
                "created_on": "2018-04-11T18:24:30.123456Z"
            }
        }
    }
}
```
</div>

</div>
