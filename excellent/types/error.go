package types

import (
	"fmt"

	"github.com/nyaruka/goflow/utils"
)

// XError is an error
type XError interface {
	error
	XPrimitive
	Equals(XError) bool
}

type xerror struct {
	native error
}

// NewXError creates a new XError
func NewXError(err error) XError {
	return xerror{native: err}
}

// NewXErrorf creates a new XError
func NewXErrorf(format string, a ...interface{}) XError {
	return NewXError(fmt.Errorf(format, a...))
}

// NewXResolveError creates a new XError when a key can't be resolved on an XResolvable
func NewXResolveError(resolvable XResolvable, key string) XError {
	val, _ := resolvable.(XValue)
	return NewXError(fmt.Errorf("%s has no property '%s'", Describe(val), key))
}

// Describe returns a representation of this type for error messages
func (x xerror) Describe() string { return "error" }

// Reduce returns the primitive version of this type (i.e. itself)
func (x xerror) Reduce(env utils.Environment) XPrimitive { return x }

// ToXText converts this type to text
func (x xerror) ToXText(env utils.Environment) XText { return NewXText(x.Native().Error()) }

// ToXBoolean converts this type to a bool
func (x xerror) ToXBoolean(env utils.Environment) XBoolean { return XBooleanFalse }

// ToXJSON is called when this type is passed to @(json(...))
func (x xerror) ToXJSON(env utils.Environment) XText { return MustMarshalToXText(x.Native().Error()) }

// MarshalJSON converts this type to internal JSON
func (x xerror) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// Native returns the native value of this type
func (x xerror) Native() error { return x.native }

func (x xerror) Error() string { return x.Native().Error() }

func (x xerror) String() string { return x.Native().Error() }

// Equals determines equality for this type
func (x xerror) Equals(other XError) bool {
	return x.String() == other.String()
}

// NilXError is the nil error value
var NilXError = NewXError(nil)
var _ XError = NilXError

// IsXError returns whether the given value is an error value
func IsXError(x XValue) bool {
	_, isError := x.(XError)
	return isError
}
