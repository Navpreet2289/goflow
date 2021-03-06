# Templates

Some properties of entities in the flow specification are _templates_ - that is their values are dynamic and are evaluated at runtime. 
Templates can contain single variables or more complex expressions. A single variable is embedded using the `@` character. For example 
the template `Hi @contact.name` contains a single variable which at runtime will be replaced with the name of the current contact.

More complex expressions can be embedded using the `@(...)` syntax. For example the template `Hi @("Dr " & upper(contact.name))` takes 
the contact name, converts it to uppercase, and the prefixes it with another string.

The `@` symbol can be escaped in templates by repeating it, ie, `Hi @@twitter` would output `Hi @twitter`.

# Context

The context is all the variables which are accessible in expressions and contains the following top-level variables:

 * `contact` the [contact](#context:contact) of the current flow run
 * `run` the current [run](#context:run)
 * `parent` the parent of the current [run](#context:run), i.e. the run that started the current run
 * `child` the child of the current [run](#context:run), i.e. the last subflow
 * `trigger` the [trigger](#context:trigger) that initiated this session

The following types appear in the context:

 * [Channel](#context:channel)
 * [Contact](#context:contact)
 * [Flow](#context:flow)
 * [Group](#context:group)
 * [Input](#context:input)
 * [Result](#context:result)
 * [Run](#context:run)
 * [Trigger](#context:trigger)
 * [URN](#context:urn)
 * [Webhook](#context:webhook)

<div class="context">
<a name="context:attachment"></a>

## Attachment

Is a media attachment on a message, and it has the following properties which can be accessed:

 * `content_type` the MIME type of the attachment
 * `url` the URL of the attachment

Examples:


```objectivec
@run.input.attachments.0.content_type → image/jpeg
@run.input.attachments.0.url → http://s3.amazon.com/bucket/test.jpg
@(json(run.input.attachments.0)) → {"content_type":"image/jpeg","url":"http://s3.amazon.com/bucket/test.jpg"}
```

<a name="context:channel"></a>

## Channel

Represents a means for sending and receiving input during a flow run. It renders as its name in a template,
and has the following properties which can be accessed:

 * `uuid` the UUID of the channel
 * `name` the name of the channel
 * `address` the address of the channel

Examples:


```objectivec
@contact.channel → My Android Phone
@contact.channel.name → My Android Phone
@contact.channel.address → +12345671111
@run.input.channel.uuid → 57f1078f-88aa-46f4-a59a-948a5739c03d
@(json(contact.channel)) → {"address":"+12345671111","name":"My Android Phone","uuid":"57f1078f-88aa-46f4-a59a-948a5739c03d"}
```

<a name="context:contact"></a>

## Contact

Represents a person who is interacting with the flow. It renders as the person's name
(or perferred URN if name isn't set) in a template, and has the following properties which can be accessed:

 * `uuid` the UUID of the contact
 * `name` the full name of the contact
 * `first_name` the first name of the contact
 * `language` the [ISO-639-3](http://www-01.sil.org/iso639-3/) language code of the contact
 * `timezone` the timezone name of the contact
 * `created_on` the datetime when the contact was created
 * `urns` all [URNs](#context:urn) the contact has set
 * `urns.[scheme]` all the [URNs](#context:urn) the contact has set for the particular URN scheme
 * `urn` shorthand for `@(format_urn(c.urns.0))`, i.e. the contact's preferred [URN](#context:urn) in friendly formatting
 * `groups` all the [groups](#context:group) that the contact belongs to
 * `fields` all the custom contact fields the contact has set
 * `fields.[snaked_field_name]` the value of the specific field
 * `channel` shorthand for `contact.urns.0.channel`, i.e. the [channel](#context:channel) of the contact's preferred URN

Examples:


```objectivec
@contact → Ryan Lewis
@contact.name → Ryan Lewis
@contact.first_name → Ryan
@contact.language → eng
@contact.timezone → America/Guayaquil
@contact.created_on → 2018-06-20T11:40:30.123456Z
@contact.urns → ["tel:+12065551212?channel=57f1078f-88aa-46f4-a59a-948a5739c03d","twitterid:54784326227#nyaruka","mailto:foo@bar.com"]
@contact.urns.0 → tel:+12065551212?channel=57f1078f-88aa-46f4-a59a-948a5739c03d
@contact.urns.tel → ["tel:+12065551212?channel=57f1078f-88aa-46f4-a59a-948a5739c03d"]
@contact.urns.mailto.0 → mailto:foo@bar.com
@contact.urn → (206) 555-1212
@contact.groups → ["Testers","Males"]
@contact.fields → {"activation_token":"AACC55","age":"23","gender":"Male","join_date":"2017-12-02T00:00:00.000000-02:00"}
@contact.fields.activation_token → AACC55
@contact.fields.gender → Male
```

<a name="context:flow"></a>

## Flow

Describes the ordered logic of actions and routers. It renders as its name in a template, and has the following
properties which can be accessed:

 * `uuid` the UUID of the flow
 * `name` the name of the flow
 * `revision` the revision number of the flow

Examples:


```objectivec
@run.flow → Registration
@child.flow → Collect Age
@run.flow.uuid → 50c3706e-fedb-42c0-8eab-dda3335714b7
@(json(run.flow)) → {"name":"Registration","revision":123,"uuid":"50c3706e-fedb-42c0-8eab-dda3335714b7"}
```

<a name="context:group"></a>

## Group

Represents a grouping of contacts. It can be static (contacts are added and removed manually through
[actions](#action:add_contact_groups)) or dynamic (contacts are added automatically by a query). It renders as its name in a
template, and has the following properties which can be accessed:

 * `uuid` the UUID of the group
 * `name` the name of the group

Examples:


```objectivec
@contact.groups → ["Testers","Males"]
@contact.groups.0.uuid → b7cf0d83-f1c9-411c-96fd-c511a4cfa86d
@contact.groups.1.name → Males
@(json(contact.groups.1)) → {"name":"Males","uuid":"4f1f98fc-27a7-4a69-bbdb-24744ba739a9"}
```

<a name="context:input"></a>

## Input

Describes input from the contact and currently we only support one type of input: `msg`. Any input has the following
properties which can be accessed:

 * `uuid` the UUID of the input
 * `type` the type of the input, e.g. `msg`
 * `channel` the [channel](#context:channel) that the input was received on
 * `created_on` the time when the input was created

An input of type `msg` renders as its text and attachments in a template, and has the following additional properties:

 * `text` the text of the message
 * `attachments` any [attachments](#context:attachment) on the message
 * `urn` the [URN](#context:urn) that the input was received on

Examples:


```objectivec
@run.input → Hi there\nhttp://s3.amazon.com/bucket/test.jpg\nhttp://s3.amazon.com/bucket/test.mp3
@run.input.type → msg
@run.input.text → Hi there
@run.input.attachments → ["http://s3.amazon.com/bucket/test.jpg","http://s3.amazon.com/bucket/test.mp3"]
@(json(run.input)) → {"attachments":[{"content_type":"image/jpeg","url":"http://s3.amazon.com/bucket/test.jpg"},{"content_type":"audio/mp3","url":"http://s3.amazon.com/bucket/test.mp3"}],"channel":{"address":"+12345671111","name":"My Android Phone","uuid":"57f1078f-88aa-46f4-a59a-948a5739c03d"},"created_on":"2000-01-01T00:00:00.000000Z","text":"Hi there","type":"msg","urn":{"display":"","path":"+12065551212","scheme":"tel"},"uuid":"9bf91c2b-ce58-4cef-aacc-281e03f69ab5"}
```

<a name="context:result"></a>

## Result

Describes a value captured during a run's execution. It might have been implicitly created by a router, or explicitly
created by a [set_run_result](#action:set_run_result) action.It renders as its value in a template, and has the following
properties which can be accessed:

 * `value` the value of the result
 * `category` the category of the result
 * `category_localized` the localized category of the result
 * `input` the input associated with the result
 * `node_uuid` the UUID of the node where the result was created
 * `created_on` the time when the result was created

Examples:


```objectivec
@run.results.favorite_color → red
@run.results.favorite_color.value → red
@run.results.favorite_color.category → Red
```

<a name="context:run"></a>

## Run

Is a single contact's journey through a flow. It records the path they have taken, and the results that have been
collected. It has several properties which can be accessed in expressions:

 * `uuid` the UUID of the run
 * `flow` the [flow](#context:flow) of the run
 * `contact` the [contact](#context:contact) of the flow run
 * `input` the [input](#context:input) of the current run
 * `results` the results that have been saved for this run
 * `results.[snaked_result_name]` the value of the specific result, e.g. `run.results.age`
 * `webhook` the last [webhook](#context:webhook) call made in the current run

Examples:


```objectivec
@run.flow.name → Registration
```

<a name="context:trigger"></a>

## Trigger

Represents something which can initiate a session with the flow engine. It has several properties which can be
accessed in expressions:

 * `type` the type of the trigger, one of "manual" or "flow"
 * `params` the parameters passed to the trigger

Examples:


```objectivec
@trigger.type → flow_action
@trigger.params → {"source": "website","address": {"state": "WA"}}
@(json(trigger)) → {"params":{"source":"website","address":{"state":"WA"}},"type":"flow_action"}
```

<a name="context:urn"></a>

## Urn

Represents a destination for an outgoing message or a source of an incoming message. It is string composed of 3
components: scheme, path, and display (optional). For example:

 - _tel:+16303524567_
 - _twitterid:54784326227#nyaruka_
 - _telegram:34642632786#bobby_

It has several properties which can be accessed in expressions:

 * `scheme` the scheme of the URN, e.g. "tel", "twitter"
 * `path` the path of the URN, e.g. "+16303524567"
 * `display` the display portion of the URN, e.g. "+16303524567"
 * `channel` the preferred [channel](#context:channel) of the URN

To render a URN in a human friendly format, use the [format_urn](expressions.html#function:format_urn) function.

Examples:


```objectivec
@contact.urns.0 → tel:+12065551212?channel=57f1078f-88aa-46f4-a59a-948a5739c03d
@contact.urns.0.scheme → tel
@contact.urns.0.path → +12065551212
@contact.urns.1.display → nyaruka
@(format_urn(contact.urns.0)) → (206) 555-1212
@(json(contact.urns.0)) → {"display":"","path":"+12065551212","scheme":"tel"}
```

<a name="context:webhook"></a>

## Webhook

Describes a call made to an external service. It has several properties which can be accessed in expressions:

 * `status` the status of the webhook - one of "success", "connection_error" or "response_error"
 * `status_code` the status code of the response
 * `body` the body of the response
 * `json` the parsed JSON response (if response body was JSON)
 * `json.[key]` sub-elements of the parsed JSON response
 * `request` the raw request made, including headers
 * `response` the raw response received, including headers

Examples:


```objectivec
@run.webhook.status_code → 200
@run.webhook.json.results.0.state → WA
```


</div>

# Functions

Templates also have access to a set of functions which can be used to further manipulate the context. Functions are called 
using the `@(function_name(args..))` syntax. For example, to title-case a contact's name in a message, you can use `@(title(contact.name))`. 
Context variables referred to within functions do not need a leading `@`. Functions can also use literal numbers or strings as arguments, for example
`@(length(split("1 2 3", " "))`.

<div class="functions">
<a name="function:abs"></a>

## abs(num)

Returns the absolute value of `num`


```objectivec
@(abs(-10)) → 10
@(abs(10.5)) → 10.5
@(abs("foo")) → ERROR
```

<a name="function:and"></a>

## and(tests...)

Returns whether all the passed in arguments are truthy


```objectivec
@(and(true)) → true
@(and(true, false, true)) → false
```

<a name="function:array"></a>

## array(values...)

Takes a list of `values` and returns them as an array


```objectivec
@(array("a", "b", 356)[1]) → b
@(join(array("a", "b", "c"), "|")) → a|b|c
@(length(array())) → 0
@(length(array("a", "b"))) → 2
```

<a name="function:boolean"></a>

## boolean(value)

Tries to convert `value` to a boolean. An error is returned if the value can't be converted.


```objectivec
@(boolean(array(1, 2))) → true
@(boolean("FALSE")) → false
@(boolean(1 / 0)) → ERROR
```

<a name="function:char"></a>

## char(num)

Returns the rune for the passed in codepoint, `num`, which may be unicode, this is the reverse of code


```objectivec
@(char(33)) → !
@(char(128512)) → 😀
@(char("foo")) → ERROR
```

<a name="function:clean"></a>

## clean(text)

Strips any non-printable characters from `text`


```objectivec
@(clean("😃 Hello \nwo\tr\rld")) → 😃 Hello world
@(clean(123)) → 123
```

<a name="function:code"></a>

## code(text)

Returns the numeric code for the first character in `text`, it is the inverse of char


```objectivec
@(code("a")) → 97
@(code("abc")) → 97
@(code("😀")) → 128512
@(code("15")) → 49
@(code(15)) → 49
@(code("")) → ERROR
```

<a name="function:datetime"></a>

## datetime(text)

Turns `text` into a date according to the environment's settings. It will return an error
if it is unable to convert the text to a date.


```objectivec
@(datetime("1979-07-18")) → 1979-07-18T00:00:00.000000-05:00
@(datetime("1979-07-18T10:30:45.123456Z")) → 1979-07-18T10:30:45.123456Z
@(datetime("2010 05 10")) → 2010-05-10T00:00:00.000000-05:00
@(datetime("NOT DATE")) → ERROR
```

<a name="function:datetime_add"></a>

## datetime_add(date, offset, unit)

Calculates the date value arrived at by adding `offset` number of `unit` to the `date`

Valid durations are "Y" for years, "M" for months, "W" for weeks, "D" for days, "h" for hour,
"m" for minutes, "s" for seconds


```objectivec
@(datetime_add("2017-01-15", 5, "D")) → 2017-01-20T00:00:00.000000-05:00
@(datetime_add("2017-01-15 10:45", 30, "m")) → 2017-01-15T11:15:00.000000-05:00
```

<a name="function:datetime_diff"></a>

## datetime_diff(date1, date2, unit)

Returns the integer duration between `date1` and `date2` in the `unit` specified.

Valid durations are "Y" for years, "M" for months, "W" for weeks, "D" for days, "h" for hour,
"m" for minutes, "s" for seconds


```objectivec
@(datetime_diff("2017-01-17", "2017-01-15", "D")) → 2
@(datetime_diff("2017-01-17 10:50", "2017-01-17 12:30", "h")) → -1
@(datetime_diff("2017-01-17", "2015-12-17", "Y")) → 2
```

<a name="function:datetime_from_parts"></a>

## datetime_from_parts(year, month, day)

Converts the passed in `year`, `month` and `day`


```objectivec
@(datetime_from_parts(2017, 1, 15)) → 2017-01-15T00:00:00.000000-05:00
@(datetime_from_parts(2017, 2, 31)) → 2017-03-03T00:00:00.000000-05:00
@(datetime_from_parts(2017, 13, 15)) → ERROR
```

<a name="function:default"></a>

## default(test, default)

Takes two arguments, returning `test` if not an error or nil or empty text, otherwise returning `default`


```objectivec
@(default(undeclared.var, "default_value")) → default_value
@(default("10", "20")) → 10
@(default("", "value")) → value
@(default(array(1, 2), "value")) → ["1","2"]
@(default(array(), "value")) → value
@(default(datetime("invalid-date"), "today")) → today
```

<a name="function:epoch"></a>

## epoch(date)

Converts `date` to the number of seconds since January 1st, 1970 GMT


```objectivec
@(epoch("2017-06-12T16:56:59.000000Z")) → 1497286619
@(epoch("2017-06-12T18:56:59.000000+02:00")) → 1497286619
@(epoch("2017-06-12T16:56:59.123456Z")) → 1497286619.123456
@(round_down(epoch("2017-06-12T16:56:59.123456Z"))) → 1497286619
```

<a name="function:field"></a>

## field(text, offset, delimiter)

Splits `text` based on the passed in `delimiter` and returns the field at `offset`.  When splitting
with a space, the delimiter is considered to be all whitespace.  (first field is 0)


```objectivec
@(field("a,b,c", 1, ",")) → b
@(field("a,,b,c", 1, ",")) →
@(field("a   b c", 1, " ")) → b
@(field("a		b	c	d", 1, "	")) →
@(field("a\t\tb\tc\td", 1, " ")) →
@(field("a,b,c", "foo", ",")) → ERROR
```

<a name="function:format_date"></a>

## format_date(date, [,format])

Turns `date` into text according to the `format` specified.

The format string can consist of the following characters. The characters
' ', ':', ',', 'T', '-' and '_' are ignored. Any other character is an error.

* `YY`        - last two digits of year 0-99
* `YYYY`      - four digits of year 0000-9999
* `M`         - month 1-12
* `MM`        - month 01-12
* `D`         - day of month, 1-31
* `DD`        - day of month, zero padded 0-31


```objectivec
@(format_date("1979-07-18T15:00:00.000000Z")) → 1979-07-18
@(format_date("1979-07-18T15:00:00.000000Z", "YYYY-MM-DD")) → 1979-07-18
@(format_date("2010-05-10T19:50:00.000000Z", "YYYY M DD")) → 2010 5 10
@(format_date("1979-07-18T15:00:00.000000Z", "YYYY")) → 1979
@(format_date("1979-07-18T15:00:00.000000Z", "M")) → 7
@(format_date("NOT DATE", "YYYY-MM-DD")) → ERROR
```

<a name="function:format_datetime"></a>

## format_datetime(date [,format [,timezone]])

Turns `date` into text according to the `format` specified and in
the optional `timezone`.

The format string can consist of the following characters. The characters
' ', ':', ',', 'T', '-' and '_' are ignored. Any other character is an error.

* `YY`        - last two digits of year 0-99
* `YYYY`      - four digits of year 0000-9999
* `M`         - month 1-12
* `MM`        - month 01-12
* `D`         - day of month, 1-31
* `DD`        - day of month, zero padded 0-31
* `h`         - hour of the day 1-12
* `hh`        - hour of the day 01-12
* `tt`        - twenty four hour of the day 01-23
* `m`         - minute 0-59
* `mm`        - minute 00-59
* `s`         - second 0-59
* `ss`        - second 00-59
* `fff`       - milliseconds
* `ffffff`    - microseconds
* `fffffffff` - nanoseconds
* `aa`        - am or pm
* `AA`        - AM or PM
* `Z`         - hour and minute offset from UTC, or Z for UTC
* `ZZZ`       - hour and minute offset from UTC

Timezone should be a location name as specified in the IANA Time Zone database, such
as "America/Guayaquil" or "America/Los_Angeles". If not specified the timezone of your
environment will be used. An error will be returned if the timezone is not recognized.


```objectivec
@(format_datetime("1979-07-18T15:00:00.000000Z")) → 1979-07-18 10:00
@(format_datetime("1979-07-18T15:00:00.000000Z", "YYYY-MM-DD")) → 1979-07-18
@(format_datetime("2010-05-10T19:50:00.000000Z", "YYYY M DD tt:mm")) → 2010 5 10 14:50
@(format_datetime("2010-05-10T19:50:00.000000Z", "YYYY-MM-DD tt:mm AA", "America/Los_Angeles")) → 2010-05-10 12:50 PM
@(format_datetime("1979-07-18T15:00:00.000000Z", "YYYY")) → 1979
@(format_datetime("1979-07-18T15:00:00.000000Z", "M")) → 7
@(format_datetime("NOT DATE", "YYYY-MM-DD")) → ERROR
```

<a name="function:format_location"></a>

## format_location(location)

Formats the given location as its name


```objectivec
@(format_location("Rwanda")) → Rwanda
@(format_location("Rwanda > Kigali")) → Kigali
```

<a name="function:format_number"></a>

## format_number(num, places, commas)

Returns `num` formatted with the passed in number of decimal `places` and optional `commas` dividing thousands separators


```objectivec
@(format_number(31337)) → 31,337.00
@(format_number(31337, 2)) → 31,337.00
@(format_number(31337, 2, true)) → 31,337.00
@(format_number(31337, 0, false)) → 31337
@(format_number("foo", 2, false)) → ERROR
```

<a name="function:format_urn"></a>

## format_urn(urn)

Turns `urn` into human friendly text


```objectivec
@(format_urn("tel:+250781234567")) → 0781 234 567
@(format_urn("twitter:134252511151#billy_bob")) → billy_bob
@(format_urn(contact.urns)) → (206) 555-1212
@(format_urn(contact.urns.2)) → foo@bar.com
@(format_urn(contact.urns.mailto)) → foo@bar.com
@(format_urn(contact.urns.mailto.0)) → foo@bar.com
@(format_urn(contact.urns.telegram)) →
@(format_urn("NOT URN")) → ERROR
```

<a name="function:from_epoch"></a>

## from_epoch(num)

Returns a new date created from `num` which represents number of nanoseconds since January 1st, 1970 GMT


```objectivec
@(from_epoch(1497286619000000000)) → 2017-06-12T11:56:59.000000-05:00
```

<a name="function:if"></a>

## if(test, true_value, false_value)

Evaluates the `test` argument, and if truthy returns `true_value`, if not returning `false_value`

If the first argument is an error that error is returned


```objectivec
@(if(1 = 1, "foo", "bar")) → foo
@(if("foo" > "bar", "foo", "bar")) → ERROR
```

<a name="function:join"></a>

## join(array, delimiter)

Joins the passed in `array` of strings with the passed in `delimiter`


```objectivec
@(join(array("a", "b", "c"), "|")) → a|b|c
@(join(split("a.b.c", "."), " ")) → a b c
```

<a name="function:json"></a>

## json(value)

Tries to return a JSON representation of `value`. An error is returned if there is
no JSON representation of that object.


```objectivec
@(json("string")) → "string"
@(json(10)) → 10
@(json(contact.uuid)) → "5d76d86b-3bb9-4d5a-b822-c9d86f5d8e4f"
```

<a name="function:left"></a>

## left(text, count)

Returns the `count` most left characters of the passed in `text`


```objectivec
@(left("hello", 2)) → he
@(left("hello", 7)) → hello
@(left("😀😃😄😁", 2)) → 😀😃
@(left("hello", -1)) → ERROR
```

<a name="function:length"></a>

## length(value)

Returns the length of the passed in text or array.

length will return an error if it is passed an item which doesn't have length.


```objectivec
@(length("Hello")) → 5
@(length("😀😃😄😁")) → 4
@(length(array())) → 0
@(length(array("a", "b", "c"))) → 3
@(length(1234)) → ERROR
```

<a name="function:lower"></a>

## lower(text)

Lowercases the passed in `text`


```objectivec
@(lower("HellO")) → hello
@(lower("hello")) → hello
@(lower("123")) → 123
@(lower("😀")) → 😀
```

<a name="function:max"></a>

## max(values...)

Takes a list of `values` and returns the greatest of them


```objectivec
@(max(1, 2)) → 2
@(max(1, -1, 10)) → 10
@(max(1, 10, "foo")) → ERROR
```

<a name="function:mean"></a>

## mean(values)

Takes a list of `values` and returns the arithmetic mean of them


```objectivec
@(mean(1, 2)) → 1.5
@(mean(1, 2, 6)) → 3
@(mean(1, "foo")) → ERROR
```

<a name="function:min"></a>

## min(values)

Takes a list of `values` and returns the smallest of them


```objectivec
@(min(1, 2)) → 1
@(min(2, 2, -10)) → -10
@(min(1, 2, "foo")) → ERROR
```

<a name="function:mod"></a>

## mod(dividend, divisor)

Returns the remainder of the division of `divident` by `divisor`


```objectivec
@(mod(5, 2)) → 1
@(mod(4, 2)) → 0
@(mod(5, "foo")) → ERROR
```

<a name="function:now"></a>

## now()

Returns the current date and time in the environment timezone


```objectivec
@(now()) → 2018-04-11T13:24:30.123456-05:00
```

<a name="function:number"></a>

## number(value)

Tries to convert `value` to a number. An error is returned if the value can't be converted.


```objectivec
@(number(10)) → 10
@(number("123.45000")) → 123.45
@(number("what?")) → ERROR
```

<a name="function:or"></a>

## or(tests...)

Returns whether if any of the passed in arguments are truthy


```objectivec
@(or(true)) → true
@(or(true, false, true)) → true
```

<a name="function:parse_datetime"></a>

## parse_datetime(text, format [,timezone])

Turns `text` into a date according to the `format` and optional `timezone` specified

The format string can consist of the following characters. The characters
' ', ':', ',', 'T', '-' and '_' are ignored. Any other character is an error.

* `YY`        - last two digits of year 0-99
* `YYYY`      - four digits of year 0000-9999
* `M`         - month 1-12
* `MM`        - month 01-12
* `D`         - day of month, 1-31
* `DD`        - day of month, zero padded 0-31
* `h`         - hour of the day 1-12
* `hh`        - hour of the day 01-12
* `tt`        - twenty four hour of the day 01-23
* `m`         - minute 0-59
* `mm`        - minute 00-59
* `s`         - second 0-59
* `ss`        - second 00-59
* `fff`       - milliseconds
* `ffffff`    - microseconds
* `fffffffff` - nanoseconds
* `aa`        - am or pm
* `AA`        - AM or PM
* `Z`         - hour and minute offset from UTC, or Z for UTC
* `ZZZ`       - hour and minute offset from UTC

Timezone should be a location name as specified in the IANA Time Zone database, such
as "America/Guayaquil" or "America/Los_Angeles". If not specified the timezone of your
environment will be used. An error will be returned if the timezone is not recognized.

Note that fractional seconds will be parsed even without an explicit format identifier.
You should only specify fractional seconds when you want to assert the number of places
in the input format.

parse_datetime will return an error if it is unable to convert the text to a datetime.


```objectivec
@(parse_datetime("1979-07-18", "YYYY-MM-DD")) → 1979-07-18T00:00:00.000000-05:00
@(parse_datetime("2010 5 10", "YYYY M DD")) → 2010-05-10T00:00:00.000000-05:00
@(parse_datetime("2010 5 10 12:50", "YYYY M DD tt:mm", "America/Los_Angeles")) → 2010-05-10T12:50:00.000000-07:00
@(parse_datetime("NOT DATE", "YYYY-MM-DD")) → ERROR
```

<a name="function:parse_json"></a>

## parse_json(text)

Tries to parse `text` as JSON, returning a fragment you can index into

If the passed in value is not JSON, then an error is returned


```objectivec
@(parse_json("[1,2,3,4]").2) → 3
@(parse_json("invalid json")) → ERROR
```

<a name="function:percent"></a>

## percent(num)

Converts `num` to text represented as a percentage


```objectivec
@(percent(0.54234)) → 54%
@(percent(1.2)) → 120%
@(percent("foo")) → ERROR
```

<a name="function:rand"></a>

## rand()

Returns a single random number between [0.0-1.0).


```objectivec
@(rand()) → 0.3849275689214193274523267973563633859157562255859375
@(rand()) → 0.607552015674623913099594574305228888988494873046875
```

<a name="function:rand_between"></a>

## rand_between()

A single random integer in the given inclusive range.


```objectivec
@(rand_between(1, 10)) → 5
@(rand_between(1, 10)) → 10
```

<a name="function:read_chars"></a>

## read_chars(text)

Converts `text` into something that can be read by IVR systems

ReadChars will split the numbers such as they are easier to understand. This includes
splitting in 3s or 4s if appropriate.


```objectivec
@(read_chars("1234")) → 1 2 3 4
@(read_chars("abc")) → a b c
@(read_chars("abcdef")) → a b c , d e f
```

<a name="function:remove_first_word"></a>

## remove_first_word(text)

Removes the 1st word of `text`


```objectivec
@(remove_first_word("foo bar")) → bar
```

<a name="function:repeat"></a>

## repeat(text, count)

Return `text` repeated `count` number of times


```objectivec
@(repeat("*", 8)) → ********
@(repeat("*", "foo")) → ERROR
```

<a name="function:replace"></a>

## replace(text, needle, replacement)

Replaces all occurrences of `needle` with `replacement` in `text`


```objectivec
@(replace("foo bar", "foo", "zap")) → zap bar
@(replace("foo bar", "baz", "zap")) → foo bar
```

<a name="function:right"></a>

## right(text, count)

Returns the `count` most right characters of the passed in `text`


```objectivec
@(right("hello", 2)) → lo
@(right("hello", 7)) → hello
@(right("😀😃😄😁", 2)) → 😄😁
@(right("hello", -1)) → ERROR
```

<a name="function:round"></a>

## round(num [,places])

Rounds `num` to the nearest value. You can optionally pass in the number of decimal places to round to as `places`.

If places < 0, it will round the integer part to the nearest 10^(-places).


```objectivec
@(round(12)) → 12
@(round(12.141)) → 12
@(round(12.6)) → 13
@(round(12.141, 2)) → 12.14
@(round(12.146, 2)) → 12.15
@(round(12.146, -1)) → 10
@(round("notnum", 2)) → ERROR
```

<a name="function:round_down"></a>

## round_down(num [,places])

Rounds `num` down to the nearest integer value. You can optionally pass in the number of decimal places to round to as `places`.


```objectivec
@(round_down(12)) → 12
@(round_down(12.141)) → 12
@(round_down(12.6)) → 12
@(round_down(12.141, 2)) → 12.14
@(round_down(12.146, 2)) → 12.14
@(round_down("foo")) → ERROR
```

<a name="function:round_up"></a>

## round_up(num [,places])

Rounds `num` up to the nearest integer value. You can optionally pass in the number of decimal places to round to as `places`.


```objectivec
@(round_up(12)) → 12
@(round_up(12.141)) → 13
@(round_up(12.6)) → 13
@(round_up(12.141, 2)) → 12.15
@(round_up(12.146, 2)) → 12.15
@(round_up("foo")) → ERROR
```

<a name="function:split"></a>

## split(text, delimiters)

Splits `text` based on the characters in `delimiters`

Empty values are removed from the returned list


```objectivec
@(split("a b c", " ")) → ["a","b","c"]
@(split("a", " ")) → ["a"]
@(split("abc..d", ".")) → ["abc","d"]
@(split("a.b.c.", ".")) → ["a","b","c"]
@(split("a|b,c  d", " .|,")) → ["a","b","c","d"]
```

<a name="function:text"></a>

## text(value)

Tries to convert `value` to text. An error is returned if the value can't be converted.


```objectivec
@(text(3 = 3)) → true
@(json(text(123.45))) → "123.45"
@(text(1 / 0)) → ERROR
```

<a name="function:text_compare"></a>

## text_compare(text1, text2)

Returns the comparison between the strings `text1` and `text2`.
The return value will be -1 if str1 is smaller than str2, 0 if they
are equal and 1 if str1 is greater than str2


```objectivec
@(text_compare("abc", "abc")) → 0
@(text_compare("abc", "def")) → -1
@(text_compare("zzz", "aaa")) → 1
```

<a name="function:title"></a>

## title(text)

Titlecases the passed in `text`, capitalizing each word


```objectivec
@(title("foo")) → Foo
@(title("ryan lewis")) → Ryan Lewis
@(title(123)) → 123
```

<a name="function:today"></a>

## today()

Returns the current date in the current timezone, time is set to midnight in the environment timezone


```objectivec
@(today()) → 2018-04-11T00:00:00.000000-05:00
```

<a name="function:tz"></a>

## tz(date)

Returns the timezone for `date``

If not timezone information is present in the date, then the environment's
timezone will be returned


```objectivec
@(tz("2017-01-15T02:15:18.123456Z")) → UTC
@(tz("2017-01-15 02:15:18PM")) → America/Guayaquil
@(tz("2017-01-15")) → America/Guayaquil
@(tz("foo")) → ERROR
```

<a name="function:tz_offset"></a>

## tz_offset(date)

Returns the offset for the timezone as text +/- HHMM for `date`

If no timezone information is present in the date, then the environment's
timezone offset will be returned


```objectivec
@(tz_offset("2017-01-15T02:15:18.123456Z")) → +0000
@(tz_offset("2017-01-15 02:15:18PM")) → -0500
@(tz_offset("2017-01-15")) → -0500
@(tz_offset("foo")) → ERROR
```

<a name="function:upper"></a>

## upper(text)

Uppercases all characters in the passed `text`


```objectivec
@(upper("Asdf")) → ASDF
@(upper(123)) → 123
```

<a name="function:url_encode"></a>

## url_encode(text)

URL encodes `text` for use in a URL parameter


```objectivec
@(url_encode("two words")) → two+words
@(url_encode(10)) → 10
```

<a name="function:weekday"></a>

## weekday(date)

Returns the day of the week for `date`, 0 is sunday, 1 is monday..


```objectivec
@(weekday("2017-01-15")) → 0
@(weekday("foo")) → ERROR
```

<a name="function:word"></a>

## word(text, index [,delimiters])

Returns the word at the passed in `index` for the passed in `text`. There is an optional final
parameter `delimiters` which is string of characters used to split the text into words.


```objectivec
@(word("bee cat dog", 0)) → bee
@(word("bee.cat,dog", 0)) → bee
@(word("bee.cat,dog", 1)) → cat
@(word("bee.cat,dog", 2)) → dog
@(word("bee.cat,dog", -1)) → dog
@(word("bee.cat,dog", -2)) → cat
@(word("bee.*cat,dog", 1, ".*=|")) → cat,dog
@(word("O'Grady O'Flaggerty", 1, " ")) → O'Flaggerty
```

<a name="function:word_count"></a>

## word_count(text [,delimiters])

Returns the number of words in `text`. There is an optional final parameter `delimiters`
which is string of characters used to split the text into words.


```objectivec
@(word_count("foo bar")) → 2
@(word_count(10)) → 1
@(word_count("")) → 0
@(word_count("😀😃😄😁")) → 4
@(word_count("bee.*cat,dog", ".*=|")) → 2
@(word_count("O'Grady O'Flaggerty", " ")) → 2
```

<a name="function:word_slice"></a>

## word_slice(text, start, end [,delimiters])

Extracts a substring from `text` spanning from `start` up to but not-including `end`. (first word is 0). A negative
end value means that all words after the start should be returned. There is an optional final parameter `delimiters`
which is string of characters used to split the text into words.


```objectivec
@(word_slice("bee cat dog", 0, 1)) → bee
@(word_slice("bee cat dog", 0, 2)) → bee cat
@(word_slice("bee cat dog", 1, -1)) → cat dog
@(word_slice("bee cat dog", 1)) → cat dog
@(word_slice("bee cat dog", 2, 3)) → dog
@(word_slice("bee cat dog", 3, 10)) →
@(word_slice("bee.*cat,dog", 1, -1, ".*=|,")) → cat dog
@(word_slice("O'Grady O'Flaggerty", 1, 2, " ")) → O'Flaggerty
```


</div>

# Router Tests

Router tests are a special class of functions which are used within the switch router. They are called in the same way as normal functions, but 
all return a test result object which by default evalutes to true or false, but can also be used to find the matching portion of the test by using
the `match` component of the result. The flow editor builds these expressions using UI widgets, but they can be used anywhere a normal template
function is used.

<div class="tests">
<a name="test:has_all_words"></a>

## has_all_words(text, words)

Tests whether all the `words` are contained in `text`

The words can be in any order and may appear more than once.


```objectivec
@(has_all_words("the quick brown FOX", "the fox")) → true
@(has_all_words("the quick brown FOX", "the fox").match) → the FOX
@(has_all_words("the quick brown fox", "red fox")) → false
```

<a name="test:has_any_word"></a>

## has_any_word(text, words)

Tests whether any of the `words` are contained in the `text`

Only one of the words needs to match and it may appear more than once.


```objectivec
@(has_any_word("The Quick Brown Fox", "fox quick")) → true
@(has_any_word("The Quick Brown Fox", "red fox")) → true
@(has_any_word("The Quick Brown Fox", "red fox").match) → Fox
```

<a name="test:has_beginning"></a>

## has_beginning(text, beginning)

Tests whether `text` starts with `beginning`

Both text values are trimmed of surrounding whitespace, but otherwise matching is strict
without any tokenization.


```objectivec
@(has_beginning("The Quick Brown", "the quick")) → true
@(has_beginning("The Quick Brown", "the quick").match) → The Quick
@(has_beginning("The Quick Brown", "the   quick")) → false
@(has_beginning("The Quick Brown", "quick brown")) → false
```

<a name="test:has_date"></a>

## has_date(text)

Tests whether `text` contains a date formatted according to our environment


```objectivec
@(has_date("the date is 2017-01-15")) → true
@(has_date("the date is 2017-01-15").match) → 2017-01-15T13:24:30.123456-05:00
@(has_date("there is no date here, just a year 2017")) → false
```

<a name="test:has_date_eq"></a>

## has_date_eq(text, date)

Tests whether `text` a date equal to `date`


```objectivec
@(has_date_eq("the date is 2017-01-15", "2017-01-15")) → true
@(has_date_eq("the date is 2017-01-15", "2017-01-15").match) → 2017-01-15T13:24:30.123456-05:00
@(has_date_eq("the date is 2017-01-15 15:00", "2017-01-15")) → true
@(has_date_eq("there is no date here, just a year 2017", "2017-06-01")) → false
@(has_date_eq("there is no date here, just a year 2017", "not date")) → ERROR
```

<a name="test:has_date_gt"></a>

## has_date_gt(text, min)

Tests whether `text` a date after the date `min`


```objectivec
@(has_date_gt("the date is 2017-01-15", "2017-01-01")) → true
@(has_date_gt("the date is 2017-01-15", "2017-01-01").match) → 2017-01-15T13:24:30.123456-05:00
@(has_date_gt("the date is 2017-01-15", "2017-03-15")) → false
@(has_date_gt("there is no date here, just a year 2017", "2017-06-01")) → false
@(has_date_gt("there is no date here, just a year 2017", "not date")) → ERROR
```

<a name="test:has_date_lt"></a>

## has_date_lt(text, max)

Tests whether `text` contains a date before the date `max`


```objectivec
@(has_date_lt("the date is 2017-01-15", "2017-06-01")) → true
@(has_date_lt("the date is 2017-01-15", "2017-06-01").match) → 2017-01-15T13:24:30.123456-05:00
@(has_date_lt("there is no date here, just a year 2017", "2017-06-01")) → false
@(has_date_lt("there is no date here, just a year 2017", "not date")) → ERROR
```

<a name="test:has_district"></a>

## has_district(text, state)

Tests whether a district name is contained in the `text`. If `state` is also provided
then the returned district must be within that state.


```objectivec
@(has_district("Gasabo", "Kigali")) → true
@(has_district("I live in Gasabo", "Kigali")) → true
@(has_district("I live in Gasabo", "Kigali").match) → Rwanda > Kigali City > Gasabo
@(has_district("Gasabo", "Boston")) → false
@(has_district("Gasabo")) → true
```

<a name="test:has_email"></a>

## has_email(text)

Tests whether an email is contained in `text`


```objectivec
@(has_email("my email is foo1@bar.com, please respond")) → true
@(has_email("my email is foo1@bar.com, please respond").match) → foo1@bar.com
@(has_email("my email is <foo@bar2.com>")) → true
@(has_email("i'm not sharing my email")) → false
```

<a name="test:has_group"></a>

## has_group(contact, group_uuid)

Returns whether the `contact` is part of group with the passed in UUID


```objectivec
@(has_group(contact, "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d")) → true
@(has_group(contact, "b7cf0d83-f1c9-411c-96fd-c511a4cfa86d").match) → Testers
@(has_group(contact, "97fe7029-3a15-4005-b0c7-277b884fc1d5")) → false
```

<a name="test:has_number"></a>

## has_number(text)

Tests whether `text` contains a number


```objectivec
@(has_number("the number is 42")) → true
@(has_number("the number is 42").match) → 42
@(has_number("the number is forty two")) → false
```

<a name="test:has_number_between"></a>

## has_number_between(text, min, max)

Tests whether `text` contains a number between `min` and `max` inclusive


```objectivec
@(has_number_between("the number is 42", 40, 44)) → true
@(has_number_between("the number is 42", 40, 44).match) → 42
@(has_number_between("the number is 42", 50, 60)) → false
@(has_number_between("the number is not there", 50, 60)) → false
@(has_number_between("the number is not there", "foo", 60)) → ERROR
```

<a name="test:has_number_eq"></a>

## has_number_eq(text, value)

Tests whether `text` contains a number equal to the `value`


```objectivec
@(has_number_eq("the number is 42", 42)) → true
@(has_number_eq("the number is 42", 42).match) → 42
@(has_number_eq("the number is 42", 40)) → false
@(has_number_eq("the number is not there", 40)) → false
@(has_number_eq("the number is not there", "foo")) → ERROR
```

<a name="test:has_number_gt"></a>

## has_number_gt(text, min)

Tests whether `text` contains a number greater than `min`


```objectivec
@(has_number_gt("the number is 42", 40)) → true
@(has_number_gt("the number is 42", 40).match) → 42
@(has_number_gt("the number is 42", 42)) → false
@(has_number_gt("the number is not there", 40)) → false
@(has_number_gt("the number is not there", "foo")) → ERROR
```

<a name="test:has_number_gte"></a>

## has_number_gte(text, min)

Tests whether `text` contains a number greater than or equal to `min`


```objectivec
@(has_number_gte("the number is 42", 42)) → true
@(has_number_gte("the number is 42", 42).match) → 42
@(has_number_gte("the number is 42", 45)) → false
@(has_number_gte("the number is not there", 40)) → false
@(has_number_gte("the number is not there", "foo")) → ERROR
```

<a name="test:has_number_lt"></a>

## has_number_lt(text, max)

Tests whether `text` contains a number less than `max`


```objectivec
@(has_number_lt("the number is 42", 44)) → true
@(has_number_lt("the number is 42", 44).match) → 42
@(has_number_lt("the number is 42", 40)) → false
@(has_number_lt("the number is not there", 40)) → false
@(has_number_lt("the number is not there", "foo")) → ERROR
```

<a name="test:has_number_lte"></a>

## has_number_lte(text, max)

Tests whether `text` contains a number less than or equal to `max`


```objectivec
@(has_number_lte("the number is 42", 42)) → true
@(has_number_lte("the number is 42", 44).match) → 42
@(has_number_lte("the number is 42", 40)) → false
@(has_number_lte("the number is not there", 40)) → false
@(has_number_lte("the number is not there", "foo")) → ERROR
```

<a name="test:has_only_phrase"></a>

## has_only_phrase(text, phrase)

Tests whether the `text` contains only `phrase`

The phrase must be the only text in the text to match


```objectivec
@(has_only_phrase("The Quick Brown Fox", "quick brown")) → false
@(has_only_phrase("Quick Brown", "quick brown")) → true
@(has_only_phrase("the Quick Brown fox", "")) → false
@(has_only_phrase("", "")) → true
@(has_only_phrase("Quick Brown", "quick brown").match) → Quick Brown
@(has_only_phrase("The Quick Brown Fox", "red fox")) → false
```

<a name="test:has_pattern"></a>

## has_pattern(text, pattern)

Tests whether `text` matches the regex `pattern`

Both text values are trimmed of surrounding whitespace and matching is case-insensitive.


```objectivec
@(has_pattern("Sell cheese please", "buy (\w+)")) → false
@(has_pattern("Buy cheese please", "buy (\w+)")) → true
@(has_pattern("Buy cheese please", "buy (\w+)").match) → Buy cheese
@(has_pattern("Buy cheese please", "buy (\w+)").match.groups[0]) → Buy cheese
@(has_pattern("Buy cheese please", "buy (\w+)").match.groups[1]) → cheese
```

<a name="test:has_phone"></a>

## has_phone(text, country_code)

Tests whether a phone number (in the passed in `country_code`) is contained in the `text`


```objectivec
@(has_phone("my number is 2067799294", "US")) → true
@(has_phone("my number is 206 779 9294", "US").match) → +12067799294
@(has_phone("my number is none of your business", "US")) → false
```

<a name="test:has_phrase"></a>

## has_phrase(text, phrase)

Tests whether `phrase` is contained in `text`

The words in the test phrase must appear in the same order with no other words
in between.


```objectivec
@(has_phrase("the quick brown fox", "brown fox")) → true
@(has_phrase("the Quick Brown fox", "quick fox")) → false
@(has_phrase("the Quick Brown fox", "")) → true
@(has_phrase("the.quick.brown.fox", "the quick").match) → the quick
```

<a name="test:has_state"></a>

## has_state(text)

Tests whether a state name is contained in the `text`


```objectivec
@(has_state("Kigali")) → true
@(has_state("Boston")) → false
@(has_state("¡Kigali!")) → true
@(has_state("¡Kigali!").match) → Rwanda > Kigali City
@(has_state("I live in Kigali")) → true
```

<a name="test:has_text"></a>

## has_text(text)

Tests whether there the text has any characters in it


```objectivec
@(has_text("quick brown")) → true
@(has_text("quick brown").match) → quick brown
@(has_text("")) → false
@(has_text(" \n")) → false
@(has_text(123)) → true
```

<a name="test:has_value"></a>

## has_value(value)

Returns whether `value` is non-nil and not an error

Note that `contact.fields` and `run.results` are considered dynamic, so it is not an error
to try to retrieve a value from fields or results which don't exist, rather these return an empty
value.


```objectivec
@(has_value(datetime("foo"))) → false
@(has_value(not.existing)) → false
@(has_value(contact.fields.unset)) → false
@(has_value("")) → false
@(has_value("hello")) → true
```

<a name="test:has_wait_timed_out"></a>

## has_wait_timed_out(run)

Returns whether the last wait timed out.


```objectivec
@(has_wait_timed_out(run)) → false
```

<a name="test:has_ward"></a>

## has_ward(text, district, state)

Tests whether a ward name is contained in the `text`


```objectivec
@(has_ward("Gisozi", "Gasabo", "Kigali")) → true
@(has_ward("I live in Gisozi", "Gasabo", "Kigali")) → true
@(has_ward("I live in Gisozi", "Gasabo", "Kigali").match) → Rwanda > Kigali City > Gasabo > Gisozi
@(has_ward("Gisozi", "Gasabo", "Brooklyn")) → false
@(has_ward("Gisozi", "Brooklyn", "Kigali")) → false
@(has_ward("Brooklyn", "Gasabo", "Kigali")) → false
@(has_ward("Gasabo")) → false
@(has_ward("Gisozi")) → true
```

<a name="test:has_webhook_status"></a>

## has_webhook_status(webhook, status)

Tests whether the passed in `webhook` call has the passed in `status`. If there is no
webhook set, then "success" will still match.


```objectivec
@(has_webhook_status(NULL, "success")) → true
@(has_webhook_status(run.webhook, "success")) → true
@(has_webhook_status(run.webhook, "connection_error")) → false
@(has_webhook_status(run.webhook, "success").match) → {"results":[{"state":"WA"},{"state":"IN"}]}
@(has_webhook_status("abc", "success")) → ERROR
```

<a name="test:is_error"></a>

## is_error(value)

Returns whether `value` is an error

Note that `contact.fields` and `run.results` are considered dynamic, so it is not an error
to try to retrieve a value from fields or results which don't exist, rather these return an empty
value.


```objectivec
@(is_error(datetime("foo"))) → true
@(is_error(run.not.existing)) → true
@(is_error(contact.fields.unset)) → true
@(is_error("hello")) → false
```

<a name="test:is_text_eq"></a>

## is_text_eq(text1, text2)

Returns whether two text values are equal (case sensitive). In the case that they
are, it will return the text as the match.


```objectivec
@(is_text_eq("foo", "foo")) → true
@(is_text_eq("foo", "FOO")) → false
@(is_text_eq("foo", "bar")) → false
@(is_text_eq("foo", " foo ")) → false
@(is_text_eq(run.status, "completed")) → true
@(is_text_eq(run.webhook.status, "success")) → true
@(is_text_eq(run.webhook.status, "connection_error")) → false
```


</div>