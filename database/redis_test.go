package database

import (
	"testing"
)

func TestClient(t *testing.T) {
	client := Redis()

	pong, err := client.Ping().Result()
	if err != nil {
		t.Errorf("err: %v", err)
	}

	if pong != "PONG" {
		t.Errorf("Can't receive PONG")
	}
}
