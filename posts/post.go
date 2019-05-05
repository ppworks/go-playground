package posts

import (
	"github.com/google/uuid"
)

// Post ...
type Post struct {
	ID string
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
