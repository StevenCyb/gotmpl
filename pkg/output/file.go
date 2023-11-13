package output

import (
	"io"
	"os"
)

func ToFile(filePath string, content []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.WriteString(file, string(content))

	return err
}
