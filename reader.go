package http

import "bufio"

func readLine(buf *bufio.Reader) ([]byte, error) {
	var line []byte
	for {
		lb, more, err := buf.ReadLine()
		if err != nil {
			return nil, err
		}

		if lb == nil {
			return nil, nil
		}

		line = append(line, lb...)
		if !more {
			break
		}
	}

	return line, nil
}
