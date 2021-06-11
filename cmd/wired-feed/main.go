package main

import (
	_ "github.com/aaronland/go-cloud-s3blob"
	_ "gocloud.dev/blob/fileblob"
)

import (
	"context"
	"fmt"
	"github.com/aaronland/go-http-server"	
	"github.com/aaronland/go-ftrain/wired"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sfomuseum/go-flags/flagset"
	"gocloud.dev/blob"
	"io"
	"log"
	"os"
	"net/http"
)

func main() {

	fs := flagset.NewFlagSet("ftrain")

	mode := fs.String("mode", "cli", "The operation mode for the application. Valid modes are: cli, lambda, server.")

	max_items := fs.Int("max-items", 15, "The maximum number of essays to include in the feed. If '-1' then all the essays will be included.")

	feed_type := fs.String("feed-type", "rss", "The syndication feed type to output. Valid options are: atom, rss")

	bucket_uri := fs.String("bucket-uri", "-", "A valid gocloud.dev/blob URI where the RSS feed will be written or '-'. If '-' then feed is written to STDOUT.")
	feed_name := fs.String("feed-name", "ftrain-wired.xml", "The filename of the RSS feed.")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")
	
	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "FTRAIN")

	if err != nil {
		log.Fatalf("Failed to set flags from environment variables, %v", err)
	}

	ctx := context.Background()

	switch *feed_type {
	case "atom", "rss":
		// pass
	default:
		log.Fatalf("Invalid feed type")
	}

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

	generate := func(ctx context.Context, wr io.Writer) error {

		f, err := wired.GenerateFeed(ctx, *max_items)

		if err != nil {
			return fmt.Errorf("Failed to generate WIRED feed, %v", err)
		}

		var f_enc string

		switch *feed_type {
		case "atom":
		
			f_atom, err := f.ToAtom()

			if err != nil {
				return fmt.Errorf("Failed to generate Atom feed, %v", err)
			}

			f_enc = f_atom

		case "rss":
		
			f_rss, err := f.ToRss()

			if err != nil {
				return fmt.Errorf("Failed to generate Atom feed, %v", err)
			}

			f_enc = f_rss
			
		default:
			// we've already validate feed type above
		}
		
		if err != nil {
			return fmt.Errorf("Failed to encode feed, %v", err)
		}

		_, err = wr.Write([]byte(f_enc))

		if err != nil {
			return fmt.Errorf("Failed to write feed, %v", err)
		}

		return nil
	}

	switch *mode {
	case "cli":
	
		err := generate(ctx, wr)

		if err != nil {
			log.Fatal(err)
		}

	case "lambda":

		fn := func(ctx context.Context) error {
		   return generate(ctx, wr)
		}
		
		lambda.Start(fn)

	case "server":

		fn := func(rsp http.ResponseWriter, req *http.Request) {
			ctx := req.Context()

			content_type := fmt.Sprintf("application/%s+xml", *feed_type)
			rsp.Header().Set("Content-type", content_type)
			
			err := generate(ctx, rsp)

			if err != nil {
			   http.Error(rsp, err.Error(), http.StatusInternalServerError)
			}

			return
		}

		handler := http.HandlerFunc(fn)

		mux := http.NewServeMux()
		mux.Handle("/", handler)
		
		s, err := server.NewServer(ctx, *server_uri)

		if err != nil {
		   	log.Fatalf("Failed to create new server, %v", err)
		}

		log.Printf("Listening on %s\n", s.Address())

		err = s.ListenAndServe(ctx, mux)

		if err != nil {
		   log.Fatalf("Failed to serve requests, %v", err)
		}
		
	default:
		log.Fatalf("Invalid or unsupported mode")
	}

}
