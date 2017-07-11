package flows

import (
	"encoding/json"
	"fmt"

	"time"

	"github.com/nyaruka/goflow/utils"
)

// Contact represents a single contact
type Contact struct {
	uuid              ContactUUID
	name              string
	language          utils.Language
	backdownLanguages utils.LanguageList
	timezone          *time.Location
	urns              URNList
	groups            GroupList
	fields            *Fields
}

// SetLanguage sets the language for this contact
func (c *Contact) SetLanguage(lang utils.Language) { c.language = lang }

// Language gets the language for this contact
func (c *Contact) Language() utils.Language { return c.language }

// PreferredLanguages gets all languages for this contact in order of preference
func (c *Contact) PreferredLanguages() utils.LanguageList {
	var languages utils.LanguageList

	if c.language != utils.NilLanguage {
		languages = append(languages, c.language)
	}

	return append(languages, c.backdownLanguages...)
}

func (c *Contact) SetTimezone(tz *time.Location) {
	if tz == nil {
		c.timezone = time.UTC
	} else {
		c.timezone = tz
	}
}
func (c *Contact) Timezone() *time.Location { return c.timezone }

func (c *Contact) SetName(name string) { c.name = name }
func (c *Contact) Name() string        { return c.name }

func (c *Contact) URNs() URNList     { return c.urns }
func (c *Contact) UUID() ContactUUID { return c.uuid }

func (c *Contact) Groups() GroupList { return GroupList(c.groups) }
func (c *Contact) AddGroup(uuid GroupUUID, name string) {
	c.groups = append(c.groups, &Group{uuid, name})
}
func (c *Contact) RemoveGroup(uuid GroupUUID) bool {
	for i := range c.groups {
		if c.groups[i].uuid == uuid {
			c.groups = append(c.groups[:i], c.groups[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Contact) Fields() *Fields { return c.fields }

func (c *Contact) Resolve(key string) interface{} {
	switch key {

	case "name":
		return c.name

	case "uuid":
		return c.uuid

	case "urns":
		return c.urns

	case "language":
		return c.language

	case "groups":
		return GroupList(c.groups)

	case "fields":
		return c.fields

	case "timezone":
		return c.timezone

	case "urn":
		return c.urns
	}

	return fmt.Errorf("No field '%s' on contact", key)
}

// Default returns our default value in the context
func (c *Contact) Default() interface{} {
	return c
}

// String returns our string value in the context
func (c *Contact) String() string {
	return c.name
}

var _ utils.VariableResolver = (*Contact)(nil)

type ContactReference struct {
	UUID ContactUUID `json:"uuid"    validate:"required,uuid4"`
	Name string      `json:"name"`
}

//------------------------------------------------------------------------------------------
// JSON Encoding / Decoding
//------------------------------------------------------------------------------------------

// ReadContact decodes a contact from the passed in JSON
func ReadContact(data json.RawMessage) (*Contact, error) {
	contact := &Contact{}
	err := json.Unmarshal(data, contact)
	err = utils.ValidateAllUnlessErr(err, contact)
	return contact, err
}

type contactEnvelope struct {
	UUID              ContactUUID        `json:"uuid"`
	Name              string             `json:"name"`
	Language          utils.Language     `json:"language"`
	BackdownLanguages utils.LanguageList `json:"backdown_languages,omitempty"`
	Timezone          string             `json:"timezone"`
	URNs              URNList            `json:"urns"`
	Groups            GroupList          `json:"groups"`
	Fields            *Fields            `json:"fields,omitempty"`
}

func (c *Contact) UnmarshalJSON(data []byte) error {
	var ce contactEnvelope
	var err error

	err = json.Unmarshal(data, &ce)
	if err != nil {
		return err
	}

	c.name = ce.Name
	c.uuid = ce.UUID
	c.language = ce.Language
	c.backdownLanguages = ce.BackdownLanguages
	tz, err := time.LoadLocation(ce.Timezone)
	if err != nil {
		return err
	}
	c.timezone = tz

	if ce.URNs == nil {
		c.urns = make(URNList, 0)
	} else {
		c.urns = ce.URNs
	}

	if ce.Groups == nil {
		c.groups = make(GroupList, 0)
	} else {
		c.groups = ce.Groups
	}

	if ce.Fields == nil {
		c.fields = NewFields()
	} else {
		c.fields = ce.Fields
	}

	return err
}

func (c *Contact) MarshalJSON() ([]byte, error) {
	var ce contactEnvelope

	ce.Name = c.name
	ce.UUID = c.uuid
	ce.Language = c.language
	ce.BackdownLanguages = c.backdownLanguages
	ce.Timezone = c.timezone.String()
	ce.URNs = c.urns
	ce.Groups = c.groups
	ce.Fields = c.fields

	return json.Marshal(ce)
}
