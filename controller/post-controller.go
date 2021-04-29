package controller

import (
	"clean-go/entity"
	"clean-go/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type contoller struct{}

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

var (
	postService service.PostService = service.NewPostService()
)

func NewPostController() PostController {
	return &contoller{}
}

// Get all posts route
func (*contoller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err.Error())
		return
	}

	ret, err := json.Marshal(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err.Error())
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(ret)
}

// Add a new post
func (*contoller) AddPost(response http.ResponseWriter, request *http.Request) {
	var post *entity.Post
	response.Header().Set("Content-type", "application/json")
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err.Error())
		return
	}

	err = postService.Validate(post)
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err.Error())
		return
	}

	post, err = postService.Create(post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err.Error())
		return
	}

	response.WriteHeader(http.StatusOK)

	ret, err := json.Marshal(post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(err.Error())
		// response.Write([]byte(`{error: Marshelling Json }`))
		return
	}
	response.Write(ret)
}
