package router

import "net/http"

type Router interface {
	GET(u string, f func(w http.ResponseWriter, r *http.Request))
	POST(u string, f func(w http.ResponseWriter, r *http.Request))
	// GET(w http.ResponseWriter, r *http.Request)
	SERVE(port string)
}
