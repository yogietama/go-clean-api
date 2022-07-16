package repository

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/yogie/go-clean-api/entity"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = "db_learn"
)

type postgreRepo struct{}

type data struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func NewPostgreRepository() PostRepository {
	return &postgreRepo{}
}

func (*postgreRepo) Save(post *entity.Post) (*entity.Post, error) {
	db := connectDB()
	defer db.Close()
	idString := strconv.Itoa(int(post.ID))

	querySaveData := "INSERT INTO posts VALUES (" + idString + ", '" + post.Title + "', '" + post.Text + "');"
	fmt.Println(querySaveData)
	_, err := db.Exec(querySaveData)
	if err != nil {
		log.Fatalf("Failed to add a new post : %v", err)
	}

	return post, nil
}
func (*postgreRepo) FindAll() ([]entity.Post, error) {
	db := connectDB()
	defer db.Close()

	getData := "SELECT * FROM posts"
	rows, _ := db.Query(getData)

	var posts []entity.Post
	for rows.Next() {
		p := data{}

		s := reflect.ValueOf(&p).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		err := rows.Scan(columns...)
		if err != nil {
			log.Fatal(err)
		}

		idInt, _ := strconv.Atoi(p.Id)
		posts = append(posts, entity.Post{
			ID:    int64(idInt),
			Title: p.Title,
			Text:  p.Text,
		})
	}

	return posts, nil

}

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres:/%s/%s:@%s/%s?sslmode=disable", user, password, host, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error===", err)
	}

	return db
}
