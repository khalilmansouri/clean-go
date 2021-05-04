package service

import (
	"clean-go/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "post is empty", err.Error())
}

func TestValidateEmptyTitle(t *testing.T) {
	testService := NewPostService(nil)
	post := entity.Post{ID: 1, Title: "", Text: "Hello there!"}
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "title is required", err.Error())
}

func TestValidateEmptyText(t *testing.T) {
	testService := NewPostService(nil)
	post := entity.Post{ID: 1, Title: "Here we go", Text: ""}
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "text is required", err.Error())
}
