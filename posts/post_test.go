package posts

import (
	"testing"
)

func TestPost(t *testing.T) {
	post := NewPost()

	post.Upsert()

	if post.ID == "" {
		t.Errorf("Failed to create UUID")
	}

	if post.CreatedAt == nil {
		t.Errorf("Failed to set CreatedAt")
	}

	t.Logf("%v", post)
}
