package nogo

import (
	"net/http"
)

type Param struct {
	Key   string
	Value string
}

type Params []Param

// ByName returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (ps Params) ByName(name string) string {
	for i := range ps {
		if ps[i].Key == name {
			return ps[i].Value
		}
	}
	return ""
}

type Handle func(*Ctx, Params)

type node struct {
	Path   string
	Handle Handle
}

type router struct {
	config  *config
	mapping []map[string]*node
}

func newRouter() *router {

	r := new(router)
	r.config = newConfig()

	r.mapping = make([]map[string]*node, 0, 1)

	return r
}

func (router *router) handle(method, path string, handle Handle) {
	n := new(node)
	mapping := make(map[string]*node)
	n.Path = path
	n.Handle = handle

	mapping[method] = n

	router.mapping = append(router.mapping, mapping)
}

func (router *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	m := router.mapping[0]["GET"]

	ctx := &Ctx{w, r}
	r.ParseForm()
	var ps Params

	m.Handle(ctx, ps)
}
