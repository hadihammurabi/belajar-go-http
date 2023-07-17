package http

import (
	"net"
)

func Listen(addr string, handler Handler) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		con, err := ln.Accept()
		if err != nil {
			return err
		}

		go handle(con, handler)
	}

}

func handle(con net.Conn, handler Handler) {
	req, err := readReq(con)
	if err != nil {
		return
	}

	defer con.Close()
	if err := handler.Serve(Ctx{req, newDefaultRes(con)}); err != nil {
		panic(err)
	}
}
