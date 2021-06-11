package wired

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"testing"
)

//go:embed paul-ford.html
var fs embed.FS

func TestRSSFeed(t *testing.T) {

	ctx := context.Background()

	body, err := fs.ReadFile("paul-ford.html")

	if err != nil {
		t.Fatalf("Failed to read test data, %v", err)
	}

	r := bytes.NewReader(body)

	f, err := GenerateRSSFeedWithReader(ctx, r, 15)

	if err != nil {
		t.Fatalf("Failed to generate RSS feed with reader, %v", err)
	}

	fmt.Println("FOO", f)
}
