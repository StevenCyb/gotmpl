package output

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput_ToWriter_Success(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer
	expectedContent := []byte("This is a test content.")

	assert.NoError(t, ToWriter(&buf, expectedContent))
	assert.Equal(t, expectedContent, buf.Bytes())
}
