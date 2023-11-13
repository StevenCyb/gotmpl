package input

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput_FromReader_Success(t *testing.T) {
	t.Parallel()

	expectedContent := "This is some test \ncontent."
	reader := strings.NewReader(expectedContent)
	actualContent, err := FromReader(reader)
	assert.NoError(t, err)
	assert.Equal(t, []byte(expectedContent), actualContent)
}
