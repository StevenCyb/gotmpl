package output

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput_ToFile_Success(t *testing.T) {
	t.Parallel()

	testFilePath := "test_output.txt"
	expectedContent := []byte("This is a test content.")

	t.Cleanup(func() {
		os.Remove(testFilePath)
	})

	assert.NoError(t, ToFile(testFilePath, expectedContent))

	actualContent, err := os.ReadFile(testFilePath)
	assert.NoError(t, err)
	assert.Equal(t, expectedContent, actualContent)
}
