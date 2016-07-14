// nogo project nogo.go
package nogo

import (
	"net/http"
)

func init() {

}

func Run() {
	var router router

	http.ListenAndServe(":8888", &router)
}
