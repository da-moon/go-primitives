package primitives_test

import (
	"testing"

	primitives "github.com/da-moon/go-primitives"
	assert "github.com/stretchr/testify/assert"
)

func TestAppendSliceValueSet(t *testing.T) {
	actual := new(primitives.AppendSliceValue)
	err := actual.Set("foo")
	assert.NoError(t, err)
	err = actual.Set("bar")
	assert.NoError(t, err)
	expected := []string{"foo", "bar"}
	assert.Equal(t, expected, []string(*actual))
}
