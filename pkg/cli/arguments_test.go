package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArguments(t *testing.T) {
	t.Parallel()

	t.Run("Empty_Success", func(t *testing.T) {
		args := []string{"program"}
		skip := ParseArguments(args)
		assert.False(t, skip)
		assert.Equal(t, "", Template)
		assert.Equal(t, "", Output)
		assert.Equal(t, StringArray{}, Data)
		assert.False(t, help)
	})

	t.Run("Filled_Success", func(t *testing.T) {
		args := []string{"program", "-tmpl", "input", "-out", "output", "-data", "data1", "-data", "data2"}
		skip := ParseArguments(args)
		assert.False(t, skip)
		assert.Equal(t, "input", Template)
		assert.Equal(t, "output", Output)
		assert.Equal(t, StringArray{"data1", "data2"}, Data)
		assert.False(t, help)
	})

	t.Run("Help", func(t *testing.T) {
		args := []string{"program", "-help"}
		skip := ParseArguments(args)
		assert.True(t, skip)
		assert.True(t, help)
	})
}
