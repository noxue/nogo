package nogo

import (
	"fmt"
	"net/http"
)

type router struct {
	config *config
}

func newRouter() *router {

	r := new(router)
	r.config = newConfig()

	return r
}

func (router *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
