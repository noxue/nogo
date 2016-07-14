package nogo

import (
	"fmt"
	"net/http"
)

type router struct {
	config config
}

func NewRouter() router {

	var r router
	r.config = NewConfig()

	return r
}

func (router *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
