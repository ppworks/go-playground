package posts

import (
	"testing"
)

func TestPost(t *testing.T) {
	post := NewPost()
	post.Content = "Dummy Content"
	post.Author = "Dummy Author"

	err := post.Upsert()
	if err != nil {
		t.Errorf("Failed to upsert: %v", err)
	}

	err = post.Fetch()
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
	err = post.Upsert()
	if err != nil {
		t.Errorf("Failed to upsert: %v", err)
	}

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

func TestPosts(t *testing.T) {
	data := []*Post{
		&Post{Author: "Author1", Content: "Content1"},
		&Post{Author: "Author2", Content: "Content2"},
		&Post{Author: "Author3", Content: "Content3"},
		&Post{Author: "Author4", Content: "Content4"},
	}

	for _, post := range data {
		err := post.Upsert()
		if err != nil {
			t.Errorf("Failed to upsert: %v", err)
		}
	}

	posts, err := Posts(0, 3)
	if err != nil {
		t.Errorf("Failed to fetch posts: %v", err)
	}

	for i, post := range posts {
		if post.Author != data[len(data)-1-i].Author {
			t.Errorf("Wrong order")
		}
	}

	total := Total()
	if total == 0 {
		t.Errorf("Failed to count: %v", total)
	}

	for _, post := range data {
		post.Delete()
	}
}
