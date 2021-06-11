package main

import (
	"context"
	"encoding/json"
	"github.com/aaronland/go-ftrain/wired"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	"os"
)

func main() {

	fs := flagset.NewFlagSet("ftrain")

	max_items := fs.Int("max-items", 15, "The maximum number of essays to include in the feed. If '-1' then all the essays will be included.")

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "FTRAIN")

	if err != nil {
		log.Fatalf("Failed to set flags from environment variables, %v", err)
	}

	ctx := context.Background()

	f, err := wired.GenerateRSSFeed(ctx, *max_items)

	if err != nil {
		log.Fatalf("Failed to generate WIRED RSS feed, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(f)

	if err != nil {
		log.Fatalf("Failed to encode RSS feed, %v", err)
	}
}
