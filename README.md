# go-ftrain

Go package for working with things related to Paul Ford.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/go-ftrain.svg)](https://pkg.go.dev/github.com/aaronland/go-ftrain)

## Tools

```
$> make cli
go build -mod vendor -o bin/wired-feed cmd/wired-feed/main.go
```

### wired-feed

Generate an Atom or RSS feed of Paul Ford's WIRED essays.

```
$> ./bin/wired-rss-feed -h
  -bucket-uri string
    	A valid gocloud.dev/blob URI where the RSS feed will be written or '-'. If '-' then feed is written to STDOUT. (default "-")
  -feed-name string
    	The filename of the RSS feed. (default "ftrain-wired.rss")
  -max-items int
    	The maximum number of essays to include in the feed. If '-1' then all the essays will be included. (default 15)
  -mode string
    	The operation mode for the application. Valid modes are: cli, lambda (default "cli")
```

```
$> ./bin/wired-rss-feed | jq '.items[]["title"]'

"Crypto Isn’t About Money. It’s About Fandom"
"Why Humans Are So Bad at Seeing the Future"
"My Dream of the Great Unbundling"
"So You Want to Prepare for Doomsday"
"The Secret, Essential Geography of the Office"
"Love the USPS? Join the Infrastructure Appreciation Society!"
"It's Time to Pick Classes for the 2073-74 School Year!"
"The Power and Paradox of Bad Software"
"‘Real’ Programming Is an Elitist Myth"
"The Infinite Loop of Supply Chains"
"We Are All Livestreamers Now, and Zoom Is Our Stage"
"Stones, Clocks, and What We Should Actually Leave Behind"
"How Technology Explodes the Concept of ‘Generations’"
"Why I (Still) Love Tech: In Defense of a Difficult Industry"
"Netflix and Google Books Are Blurring the Line Between Past and Present"
"Meet the Web's Operating System: HTTP"
```

## See also

* https://ftrain.com/
* https://gocloud.dev/howto/blob
* https://github.com/aaronland/go-cloud-s3blob
* https://github.com/gorilla/feeds