package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecutor_Set(t *testing.T) {
	t.Parallel()

	t.Run("SingleData_Success", func(t *testing.T) {
		t.Parallel()

		executor := New()
		executor.SetTemplate([]byte("Hello, {{.Name}}!"))
		executor.SetData([]interface{}{map[string]interface{}{"Name": "John"}})
	})

	t.Run("MultiData_Success", func(t *testing.T) {
		t.Parallel()

		executor := New()
		executor.SetTemplate([]byte("Hello, {{.Name}}!"))
		executor.SetData([]interface{}{
			map[string]interface{}{"Name": "John"},
			map[string]interface{}{"Something": "else"},
		})
	})
}

func TestExecutor_Execute(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		executor := New()
		executor.SetTemplate([]byte("Hello, {{.Name}}!"))
		executor.SetData([]interface{}{map[string]interface{}{"Name": "John"}})

		output, err := executor.Execute()
		assert.NoError(t, err)
		assert.Equal(t, []byte("Hello, John!"), output)
	})

	t.Run("InvalidTemplate_Fail", func(t *testing.T) {
		t.Parallel()

		executor := New()
		executor.SetTemplate([]byte("Hello, {{.Name!"))
		executor.SetData([]interface{}{map[string]interface{}{"Name": "John"}})

		_, err := executor.Execute()
		assert.Error(t, err)
	})
}
