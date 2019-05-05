package posts

import (
	"testing"
)

func TestPost(t *testing.T) {
	post := NewPost()
	post.Content = "Dummy Content"
	post.Author = "Dummy Author"

	post.Upsert()
	post.Fetch()

	err := post.Fetch()
	if err != nil {
		t.Errorf("Failed to fetch: %v", err)
	}

	if post.ID == "" {
		t.Errorf("Failed to create UUID")
	}

	if post.CreatedAt == nil {
		t.Errorf("Failed to set CreatedAt")
	}

	post.Content = "Updated Content"
	post.Upsert()

	err = post.Fetch()
	if err != nil {
		t.Errorf("Failed to fetch: %v", err)
	}

	if post.Content != "Updated Content" {
		t.Errorf("Failed to set Content")
	}

	t.Logf("%v", post)
	post.Delete()
}
