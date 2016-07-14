// nogo project nogo.go
package nogo

import (
	"net/http"
)

func Run() {
	http.ListenAndServe("/", &Router{})
}
