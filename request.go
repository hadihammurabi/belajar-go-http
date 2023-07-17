package http

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
	"strconv"
	"strings"
)

type Req struct {
	Method  string
	Path    string
	Headers Headers
	Body    io.ReadCloser
}

func readReq(con net.Conn) (Req, error) {
	reader := bufio.NewReader(con)
	method, path, err := readProtocol(reader)
	if err != nil {
		return Req{}, err
	}

	headers, err := readHeaders(reader)
	if err != nil {
		return Req{}, err
	}

	r := Req{
		Method:  string(method),
		Path:    string(path),
		Headers: headers,
	}

	if v, ok := r.Headers["Content-Length"]; ok {
		length, err := strconv.Atoi(v)
		if err != nil {
			return Req{}, errors.New("invalid content length")
		}

		body, err := readBody(reader, length)
		if err != nil {
			return Req{}, err
		}

		r.Body = body
	}

	return r, nil
}

func readProtocol(buf *bufio.Reader) (method []byte, path []byte, err error) {
	firstLineB, err := readLine(buf)
	if err != nil {
		return
	}

	firstLine := strings.Split(string(firstLineB), " ")
	if len(firstLine) < 3 {
		err = errors.New("invalid request")
		return
	}

	method = []byte(firstLine[0])
	path = []byte(firstLine[1])

	return
}

func readHeaders(buf *bufio.Reader) (Headers, error) {
	headers := make(Headers)
	for {
		l, err := readLine(buf)
		if err != nil {
			return nil, err
		}

		if string(l) == "" {
			break
		}

		header := strings.Split(string(l), ": ")
		if len(header) < 2 {
			return nil, errors.New("invalid header")
		}

		headers[header[0]] = header[1]
	}
	return headers, nil
}

func readBody(buf *bufio.Reader, length int) (io.ReadCloser, error) {
	body := make([]byte, length)
	_, err := io.ReadFull(buf, body)
	if err != nil {
		return nil, err
	}

	bodyC := io.NopCloser(bytes.NewReader(body))
	return bodyC, nil
}
