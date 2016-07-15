// nogo project nogo.go
package nogo

import (
	"net/http"
)

func init() {

}

func Run() {
	router := NewRouter()
	router.config.AddFile("config.conf")
	router.config.AddFile("config1.conf")
	router.config.AddFile("config2.conf")
	router.config.AddFile("config3.conf")
	router.config.AddFile("config4.conf")
	router.config.AddFile("config5.conf")
	router.config.ReadAllConfig()

	http.ListenAndServe(":8888", &router)
}
