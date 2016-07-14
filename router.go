package nogo

import (
	"fmt"
	"net/http"
)

type Router struct {
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
