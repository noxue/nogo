// nogo project nogo.go
package nogo

import (
	"net/http"
)

func init() {

}

func Run() {
	router := NewRouter()

	http.ListenAndServe(":8888", &router)
}
