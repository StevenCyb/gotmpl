package engine

import (
	"bytes"
	"fmt"

	"html/template"
)

type Executor struct {
	template []byte
	data     interface{}
}

func New() *Executor {
	return &Executor{}
}

func (e *Executor) SetTemplate(template []byte) {
	e.template = template
}

func (e *Executor) SetData(data []interface{}) {
	if len(data) == 0 {
		e.data = map[string]interface{}{}

		return
	} else if len(data) == 1 {
		e.data = data[0]

		return
	}

	e.data = data
}

func (e *Executor) Execute() ([]byte, error) {
	tmpl, err := template.New("template").Parse(string(e.template))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %v", err)
	}

	var outputBuffer bytes.Buffer
	if err = tmpl.Execute(&outputBuffer, e.data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %v", err)
	}

	return outputBuffer.Bytes(), nil
}
