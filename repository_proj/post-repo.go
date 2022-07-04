package repository_proj

import (
	"context"
	"log"

	"../entity_db"

	"cloud.google.com/go/firestore"
)

type PostRepository interface {
	Save(post *entity_db.Post) (*entity_db.Post, error)
	FindAll() ([]entity_db.Post, error)
}

type repo struct{}

const (
	projectId      = "fir-tutorial-12bf2"
	collectionName = "posts"
)

func NewPostRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity_db.Post) (*entity_db.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create a firestore client : %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to add a new post : %v", err)
		return nil, err
	}

	return post, nil

}

func (*repo) FindAll() ([]entity_db.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create a firestore client : %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity_db.Post

	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()

		if err != nil {
			// log.Fatalf("Failed to iterate the list of posts : %v", err)
			break
		}

		post := entity_db.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}
