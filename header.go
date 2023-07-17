package http

import (
	"bytes"
	"fmt"
)

type Headers map[string]string

func (h Headers) toHttpBytes() []byte {
	headerString := bytes.Buffer{}
	for k, v := range h {
		headerString.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}

	headerString.WriteString("\r\n")
	return headerString.Bytes()
}
