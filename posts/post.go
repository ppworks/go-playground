package posts

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
}

// Post ...
type Post struct {
	ID             string
	IncrementalKey int
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	Author         string
	Content        string
}

// Posts return Post slice from DB
func Posts(page int, per int) ([]*Post, error) {
	limit, offset := per, page*per
	posts := make([]*Post, per)

	rows, err := db.Query(`
		SELECT id, incremental_key, created_at, updated_at, author, content
		FROM posts
		ORDER BY incremental_key DESC
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		return posts, err
	}

	for i := 0; rows.Next(); i++ {
		post := NewPost()
		err = rows.Scan(&post.ID, &post.IncrementalKey, &post.CreatedAt, &post.UpdatedAt, &post.Author, &post.Content)
		if err != nil {
			return posts, err
		}
		posts[i] = post
	}
	return posts, nil
}

// Total returns posts count
func Total() int {
	total := 0
	db.QueryRow(`
		SELECT COUNT(id) FROM posts
	`).Scan(&total)

	return total
}

// NewPost return Post ref
func NewPost() *Post {
	p := new(Post)
	return p
}

// Fetch a post from DB
func (p *Post) Fetch() (err error) {
	err = db.QueryRow(`
		SELECT id, incremental_key, created_at, updated_at, author, content
		FROM posts
		WHERE id = $1
	`, p.ID).Scan(&p.ID, &p.IncrementalKey, &p.CreatedAt, &p.UpdatedAt, &p.Author, &p.Content)
	return
}

// Upsert post data
func (p *Post) Upsert() error {
	now := time.Now()
	if p.CreatedAt == nil {
		p.CreatedAt = &now
	}
	p.UpdatedAt = &now

	if p.ID == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			p.ID = ""
		} else {
			p.ID = id.String()
		}
	}

	sql := `
		INSERT INTO posts(id, created_at, updated_at, author, content)
		VALUES($1, $2, $3, $4, $5)
		ON CONFLICT(id)
		DO UPDATE SET
			id = $1,
			updated_at = $3,
			author = $4,
			content = $5
		RETURNING id
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.QueryRow(p.ID, p.CreatedAt, p.UpdatedAt, p.Author, p.Content)
	return nil
}

func (p *Post) Delete() {
	db.Exec(`
		DELETE FROM posts
		WHERE id = $1
	`, p.ID)
	return
}

// String for fmt.Stringer
func (p *Post) String() string {
	return fmt.Sprintf("Post{ID: %s, IncrementalKey: %d, CreatedAt: %s, UpdatedAt: %s, Author: %s, Content: %s}", p.ID, p.IncrementalKey, p.CreatedAt, p.UpdatedAt, p.Author, p.Content)
}
