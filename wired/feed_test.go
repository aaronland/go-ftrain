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

func TestFeed(t *testing.T) {

	ctx := context.Background()

	body, err := fs.ReadFile("paul-ford.html")

	if err != nil {
		t.Fatalf("Failed to read test data, %v", err)
	}

	r := bytes.NewReader(body)

	f, err := GenerateFeedWithReader(ctx, r, 15)

	if err != nil {
		t.Fatalf("Failed to generate feed with reader, %v", err)
	}

	fmt.Println("FOO", f)
}
