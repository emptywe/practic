package gorillaRouter

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func NewGorillaHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", h.Hello).Methods("GET")
	return router
}
