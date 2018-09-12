package assets

import (
	"github.com/nyaruka/goflow/utils"
)

// ChannelUUID is the UUID of a channel
type ChannelUUID utils.UUID

// NilChannelUUID is an empty channel UUID
const NilChannelUUID ChannelUUID = ChannelUUID("")

// ChannelRole is a role that a channel can perform
type ChannelRole string

// different roles that channels can perform
const (
	ChannelRoleSend    ChannelRole = "send"
	ChannelRoleReceive ChannelRole = "receive"
	ChannelRoleCall    ChannelRole = "call"
	ChannelRoleAnswer  ChannelRole = "answer"
	ChannelRoleUSSD    ChannelRole = "ussd"
)

// Channel is something that can send/receive messages
type Channel interface {
	UUID() ChannelUUID
	Name() string
	Address() string
	Schemes() []string
	Roles() []ChannelRole
	ParentUUID() ChannelUUID
	Country() string
	MatchPrefixes() []string
}

// FieldType is the data type of values for each field
type FieldType string

// field value types
const (
	FieldTypeText     FieldType = "text"
	FieldTypeNumber   FieldType = "number"
	FieldTypeDatetime FieldType = "datetime"
	FieldTypeWard     FieldType = "ward"
	FieldTypeDistrict FieldType = "district"
	FieldTypeState    FieldType = "state"
)

// Field is a custom contact property
type Field interface {
	Key() string
	Name() string
	Type() FieldType
}

// GroupUUID is the UUID of a group
type GroupUUID utils.UUID

// Group is a set of contacts
type Group interface {
	UUID() GroupUUID
	Name() string
	Query() string
}

// LabelUUID is the UUID of a label
type LabelUUID utils.UUID

// Label is something that can be applied a message
type Label interface {
	UUID() LabelUUID
	Name() string
}

// Location is a geographical entity
type Location interface {
	Name() string
	Aliases() []string
	Children() []Location
}

// Resthook is a set of URLs which are subscribed to the named event
type Resthook interface {
	Slug() string
	Subscribers() []string
}

// AssetSource is a source of assets
type AssetSource interface {
	Channels() ([]Channel, error)
	Fields() ([]Field, error)
	Groups() ([]Group, error)
	Labels() ([]Label, error)
	Locations() ([]Location, error)
	Resthooks() ([]Resthook, error)

	HasLocations() bool
}
