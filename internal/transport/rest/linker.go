package rest

import (
	"ms-hw/internal/transport/rest/handler"

	"github.com/gorilla/mux"
)

func Linker(auth *handler.AuthHandler) *mux.Router {
	router := mux.NewRouter()
	router.Path("/auth").HandlerFunc(auth.GrantAccess).Name("GrantAccess")

	return router

}
