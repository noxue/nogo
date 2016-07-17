package nogo

import (
	"net/http"
)

type Ctx struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}
