package postgres

import (
	"context"
	"strings"
	"testing"
)

func TestOpenRejectsInvalidConfiguration(t *testing.T) {
	for _, url := range []string{"", "postgres://user:secret@localhost:invalid/db"} {
		pool, err := Open(context.Background(), url)
		if pool != nil {
			pool.Close()
			t.Fatal("expected no pool")
		}
		if err == nil {
			t.Fatal("expected a configuration error")
		}
		if strings.Contains(err.Error(), "secret") {
			t.Fatal("error exposed credentials")
		}
	}
}
