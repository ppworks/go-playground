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
	ID        string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Author    string
	Content   string
}

// NewPost return Post ref
func NewPost() *Post {
	p := new(Post)
	return p
}

func (p *Post) Fetch() (err error) {
	err = db.QueryRow(`
		SELECT id, created_at, updated_at, author, content
		FROM posts
		WHERE id = $1
	`, p.ID).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.Author, &p.Content)
	return
}

// Upsert post data
func (p *Post) Upsert() {
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
		return
	}
	defer stmt.Close()
	stmt.QueryRow(p.ID, p.CreatedAt, p.UpdatedAt, p.Author, p.Content)
	return
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
	return fmt.Sprintf("Post{ID: %s, CreatedAt: %s, UpdatedAt: %s, Author: %s, Content: %s}", p.ID, p.CreatedAt, p.UpdatedAt, p.Author, p.Content)
}
