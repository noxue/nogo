// nogo project nogo.go
package nogo

import (
	"net/http"
)

type nogo struct {
	router *router
}

func (nogo *nogo) AddConfig(file string) {
	nogo.router.config.AddFile(file)
}

func NewNogo() *nogo {
	nogo := new(nogo)
	nogo.router = newRouter()

	nogo.AddConfig("config/app.conf")

	return nogo
}

func init() {

}

func (nogo *nogo) Run() {

	nogo.router.config.ReadAllConfig()

	LogPrintln("Nogo成功启动")
	err := http.ListenAndServe(nogo.router.config.getString("listen"), nogo.router)
	if err != nil {
		LogFatalln(err.Error())
	}
}
