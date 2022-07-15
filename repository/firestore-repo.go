package repository

/*

export GOOGLE_APPLICATION_CREDENTIALS="/Documents/Yogie/Tutorial/Creds/firebase-creds.json"
*/

import (
	"context"
	"log"

	"github.com/yogie/go-clean-api/entity"

	"cloud.google.com/go/firestore"
)

type repo struct{}

const (
	projectId      = "fir-tutorial-12bf2"
	collectionName = "posts"
)

func NewFirestoreRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
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

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create a firestore client : %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post

	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()

		if err != nil {
			// log.Fatalf("Failed to iterate the list of posts : %v", err)
			break
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
