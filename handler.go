package http

type Handler interface {
	Serve(Ctx) error
}
