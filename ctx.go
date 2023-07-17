package http

type Ctx struct {
	req Req
	res Res
}

func (c *Ctx) Status(status int) {
	c.res.Status = status
}

func (c *Ctx) Write(b []byte) (int, error) {
	return c.res.Write(b)
}
