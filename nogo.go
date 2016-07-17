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

func (nogo *nogo) Get(path string, handle Handle) {
	nogo.router.handle("GET", path, handle)
}

func (nogo *nogo) Run() {

	nogo.router.config.ReadAllConfig()

	var err error
	listen := nogo.router.config.getString("listen")

	proto := nogo.router.config.getString("proto")

	//default http
	if proto == "" {
		proto = "http"
	}

	if proto == "http" {
		if listen == "" {
			listen = ":80"
		}
		LogPrintln("listen：" + listen)

		err = http.ListenAndServe(listen, nogo.router)

	} else if proto == "https" {

		if listen == "" {
			listen = ":443"
		}

		certFile := nogo.router.config.getString("certfile")
		keyFile := nogo.router.config.getString("keyfile")

		if certFile == "" || keyFile == "" {
			LogFatalln("Not Found CertFile or KeyFile")
		}

		LogPrintln("listen：" + listen)
		err = http.ListenAndServeTLS(listen, certFile, keyFile, nogo.router)

	} else {
		LogFatalln("Not Suport Proto [ " + proto + " ]")
	}

	if err != nil {
		LogFatalln(err.Error())
	}
}
