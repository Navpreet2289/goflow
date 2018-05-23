package types_test

import (
	"fmt"
	"testing"

	"github.com/nyaruka/goflow/excellent/types"
	"github.com/nyaruka/goflow/utils"

	"github.com/stretchr/testify/assert"
)

func TestXError(t *testing.T) {
	env := utils.NewDefaultEnvironment()

	err1 := types.NewXError(fmt.Errorf("I failed"))
	assert.Equal(t, types.NewXText("I failed"), err1.ToXText(env))
	assert.Equal(t, types.NewXText(`"I failed"`), err1.ToXJSON(env))
	assert.Equal(t, types.XBooleanFalse, err1.ToXBoolean(env))
	assert.Equal(t, "I failed", err1.String())
	assert.Equal(t, "I failed", err1.Error())

	err2 := types.NewXResolveError(nil, "foo")
	assert.Equal(t, "null has no property 'foo'", err2.Error())
}
