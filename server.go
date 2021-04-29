package main

import (
	"clean-go/controller"
	"clean-go/repository"
	"clean-go/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepo()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	router := mux.NewRouter()
	const PORT string = ":8080"
	router.HandleFunc("/", postController.GetPosts).Methods("GET")
	router.HandleFunc("/", postController.AddPost).Methods("POST")
	log.Println("Server is listening on port ", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
