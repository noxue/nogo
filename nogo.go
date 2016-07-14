// nogo project nogo.go
package nogo

import (
	"net/http"
)

func Run() {
	var router Router
	http.ListenAndServe(":8888", &router)
}
