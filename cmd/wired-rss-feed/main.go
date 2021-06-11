package main

import (
	"context"
	"github.com/aaronland/go-ftrain/wired"
	"log"
	"flag"
	"os"
	"encoding/json"
)

func main() {

	flag.Parse()
	
	ctx := context.Background()
	
	f, err := wired.GenerateRSSFeed(ctx, -1)

	if err != nil {
		log.Fatalf("Failed to generate WIRED RSS feed, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(f)

	if err != nil {
		log.Fatalf("Failed to encode RSS feed, %v", err)
	}
}
