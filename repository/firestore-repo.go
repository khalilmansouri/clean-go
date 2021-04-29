package repository

import (
	"context"
	"log"
	"math/rand"

	"clean-go/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type repo struct{}

func NewFirestoreRepo() PostRepository {
	return &repo{}
}

const (
	projectId      string = "pragmatic-a"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
		return nil, err
	}

	post.ID = rand.Int63()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	defer client.Close()

	if err != nil {
		log.Fatalf("Failed to Add post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate doc iterator: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
