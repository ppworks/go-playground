package posts

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Post ...
type Post struct {
	ID        string
	CreatedAt *time.Time
	Content   string
	Author    string
}

// NewPost return Post ref
func NewPost() *Post {
	post := new(Post)

	id, err := uuid.NewRandom()
	if err != nil {
		post.ID = ""
	} else {
		post.ID = id.String()
	}

	return post
}

// Save post data
func (p *Post) Save() {
	now := time.Now()
	p.CreatedAt = &now
}

// String for fmt.Stringer
func (p *Post) String() string {
	return fmt.Sprintf("Post{ID: %s, CreatedAt: %s}", p.ID, p.CreatedAt)
}
