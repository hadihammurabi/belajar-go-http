package main

import "http"

type Handler struct{}

func (Handler) Serve(c http.Ctx) error {
	c.Status(404)
	c.Write([]byte("mantap\r\n"))
	return nil
}

func main() {
	http.Listen(":8080", &Handler{})
}
