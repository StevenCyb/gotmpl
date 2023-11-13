package output

import (
	"io"
)

func ToWriter(writer io.Writer, content []byte) error {
	_, err := io.WriteString(writer, string(content))
	return err
}
