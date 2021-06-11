package main

import (
	_ "github.com/aaronland/go-cloud-s3blob"
	_ "gocloud.dev/blob/fileblob"
)

import (
	// aws_events "github.com/aws/aws-lambda-go/events"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aaronland/go-ftrain/wired"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sfomuseum/go-flags/flagset"
	"gocloud.dev/blob"
	"io"
	"log"
	"os"
)

func main() {

	fs := flagset.NewFlagSet("ftrain")

	mode := fs.String("mode", "cli", "The operation mode for the application. Valid modes are: cli, lambda")

	max_items := fs.Int("max-items", 15, "The maximum number of essays to include in the feed. If '-1' then all the essays will be included.")

	bucket_uri := fs.String("bucket-uri", "-", "A valid gocloud.dev/blob URI where the RSS feed will be written or '-'. If '-' then feed is written to STDOUT.")
	feed_name := fs.String("feed-name", "ftrain-wired.rss", "The filename of the RSS feed.")

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "FTRAIN")

	if err != nil {
		log.Fatalf("Failed to set flags from environment variables, %v", err)
	}

	ctx := context.Background()

	var wr io.Writer

	switch *bucket_uri {
	case "-":
		wr = os.Stdout
	default:

		bucket, err := blob.OpenBucket(ctx, *bucket_uri)

		if err != nil {
			log.Fatalf("Failed to open bucket, %v", err)
		}

		defer bucket.Close()

		bucket_wr, err := bucket.NewWriter(ctx, *feed_name, nil)

		if err != nil {
			log.Fatalf("Failed to open writer, %v", err)
		}

		defer bucket_wr.Close()
		wr = bucket_wr
	}

	generate := func(ctx context.Context) error {

		f, err := wired.GenerateRSSFeed(ctx, *max_items)

		if err != nil {
			return fmt.Errorf("Failed to generate WIRED RSS feed, %v", err)
		}

		enc := json.NewEncoder(wr)
		err = enc.Encode(f)

		if err != nil {
			return fmt.Errorf("Failed to encode RSS feed, %v", err)
		}

		return nil
	}

	switch *mode {
	case "cli":
		err := generate(ctx)

		if err != nil {
			log.Fatal(err)
		}

	case "lambda":
		lambda.Start(generate)
	default:
		log.Fatalf("Invalid or unsupported mode")
	}

}
