package http

import (
	"fmt"
	"net"
)

type Res struct {
	con     net.Conn
	Status  int
	Headers Headers

	writedHeader bool
}

func newDefaultRes(con net.Conn) Res {
	defaultHeaders := Headers{
		"Content-Type": "text/plain",
	}
	return Res{
		con:          con,
		Status:       200,
		Headers:      defaultHeaders,
		writedHeader: false,
	}
}

func (r *Res) Write(b []byte) (int, error) {
	if !r.writedHeader {
		r.con.Write([]byte(fmt.Sprintf("HTTP/1.1 %d\r\n", r.Status)))
		r.con.Write([]byte(r.Headers.toHttpBytes()))
		r.writedHeader = true
	}

	return r.con.Write(b)
}
