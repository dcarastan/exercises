package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

// MyController struct
type MyController struct {
	AppController
	*render.Render
}

// Index func
func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

func main() {
	c := &MyController{Render: render.New(render.Options{})}
	http.ListenAndServe(":8888", c.Action(c.Index))
}
