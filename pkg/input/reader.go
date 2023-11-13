package input

import (
	"bufio"
	"io"
)

func FromReader(reader io.Reader) ([]byte, error) {
	content := ""
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	return []byte(content[:len(content)-1]), scanner.Err()
}
