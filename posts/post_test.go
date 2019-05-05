package posts

import (
	"testing"
)

func TestPost(t *testing.T) {
	post := NewPost()

	if post.ID == "" {
		t.Errorf("Failed to create UUID")
	}
}
