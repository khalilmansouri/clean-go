package controller

import (
	"clean-go/entity"
	"clean-go/service"
	"encoding/json"
	"log"
	"net/http"
)

type contoller struct{}

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

var (
	serv service.PostService = service.NewPostService()
)

func NewPostController() PostController {
	return &contoller{}
}

// Get all posts route
func (*contoller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	posts, err := serv.FindAll()
	if err != nil {
		log.Fatalf("Failed to fetch posts : %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{Failed to fetch posts}`))
		return
	}

	ret, err := json.Marshal(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{error: Marshelling Json }`))
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
		response.Write([]byte(`{error: Marshelling Json }`))
		return
	}

	post, err = serv.Create(post)

	if err != nil {
		log.Fatalf("Failed to save Post %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{error: Failed to save Post }`))
		return
	}

	response.WriteHeader(http.StatusOK)

	ret, err := json.Marshal(post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{error: Marshelling Json }`))
		return
	}
	response.Write(ret)
}
