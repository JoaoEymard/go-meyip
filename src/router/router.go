package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"../handlers"
)

type Router struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

var routersAPI = []Router{
	{
		URI:      "/myip",
		Method:   http.MethodGet,
		Function: handlers.ClientRequestIp,
	},
}

func Configuration(r *mux.Router) *mux.Router {
	for _, router := range routersAPI {
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}

	return r
}
