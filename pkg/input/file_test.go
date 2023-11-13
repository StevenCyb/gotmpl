package input

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput_FromFile(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		expectedContent := "test_content"

		tmpfile, err := os.CreateTemp("", "unit_test*.txt")
		assert.NoError(t, err)

		_, err = tmpfile.WriteString(expectedContent)
		assert.NoError(t, err)

		t.Cleanup(func() {
			tmpfile.Close()
			os.Remove(tmpfile.Name())
		})

		actualContent, err := FromFile(tmpfile.Name())
		assert.NoError(t, err)
		assert.Equal(t, []byte(expectedContent), actualContent)
	})

	t.Run("Fail", func(t *testing.T) {
		t.Parallel()

		_, err := FromFile("non_existent_file.txt")
		assert.Error(t, err)
	})
}
