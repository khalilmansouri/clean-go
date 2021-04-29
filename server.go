package main

import (
	"clean-go/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	postController controller.PostController = controller.NewPostController()
)

func main() {
	router := mux.NewRouter()
	const PORT string = ":8080"
	router.HandleFunc("/", postController.GetPosts).Methods("GET")
	router.HandleFunc("/", postController.AddPost).Methods("POST")
	log.Println("Server is listening on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
