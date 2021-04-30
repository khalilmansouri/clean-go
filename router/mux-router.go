package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(u string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(u, f).Methods("GET")
}
func (*muxRouter) POST(u string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(u, f).Methods("POST")
}
func (*muxRouter) SERVE(port string) {
	log.Println("Server is listening on port ", port)
	log.Fatal(http.ListenAndServe(port, muxDispatcher))
}
