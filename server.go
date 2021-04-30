package main

import (
	"clean-go/controller"
	"clean-go/repository"
	"clean-go/router"
	"clean-go/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepo()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	server         router.Router             = router.NewMuxRouter()
)

func main() {
	const PORT string = ":8080"
	server.GET("/", postController.GetPosts)
	server.POST("/", postController.AddPost)
	server.SERVE(PORT)
}
