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
$> ./bin/wired-feed -h
  -bucket-uri string
    	A valid gocloud.dev/blob URI where the RSS feed will be written or '-'. If '-' then feed is written to STDOUT. (default "-")
  -feed-name string
    	The filename of the RSS feed. (default "ftrain-wired.xml")
  -feed-type string
    	The syndication feed type to output. Valid options are: atom, rss (default "rss")
  -max-items int
    	The maximum number of essays to include in the feed. If '-1' then all the essays will be included. (default 15)
  -mode string
    	The operation mode for the application. Valid modes are: cli, lambda, server. (default "cli")
  -server-uri string
    	A valid aaronland/go-http-server URI. (default "http://localhost:8080")
```	

For example:

```
$> ./bin/wired-feed | xmllint --xpath "//item/title/text()" -

Crypto Isn&#x2019;t About Money. It&#x2019;s About FandomWhy Humans Are So Bad at Seeing the FutureMy Dream of the Great UnbundlingSo You Want to Prepare for DoomsdayThe Secret, Essential Geography of the OfficeLove the USPS? Join the Infrastructure Appreciation Society!It's Time to Pick Classes for the 2073-74 School Year!The Power and Paradox of Bad Software&#x2018;Real&#x2019; Programming Is an Elitist MythThe Infinite Loop of Supply ChainsWe Are All Livestreamers Now, and Zoom Is Our StageStones, Clocks, and What We Should Actually Leave BehindHow Technology Explodes the Concept of &#x2018;Generations&#x2019;Why I (Still) Love Tech: In Defense of a Difficult IndustryNetflix and Google Books Are Blurring the Line Between Past and Present
```

#### Modes

##### cli

Run the `wired-feed` tool from the command line.

##### lambda

Run `wired-feed` tool as an AWS Lambda function.

Command line options are set using environment variables. The rules for defining command line flags as environment variables are:

* Replace all instances of `-` with `_`.
* Upper-case the flag.
* Append the flag with `FTRAIN_`.

For example:

| Name | Value |
| --- | --- |
| FTRAIN_MODE | lambda |
| FTRAIN_MAX_ITEMS | 10 | 
| FTRAIN_FEED_TYPE | atom |

##### server

Run the `wired-feed` tool as an HTTP service.

For example:

```
$> ./bin/wired-feed -mode server 
2021/06/11 09:31:52 Listening on http://localhost:8080

$> curl -s localhost:8080 | xmllint --xpath "//guid" -
<guid>https://www.wired.com/story/crypto-isnt-about-money-its-about-fandom/</guid><guid>https://www.wired.com/story/why-humans-are-so-bad-at-seeing-the-future/</guid><guid>https://www.wired.com/story/my-dream-of-the-great-unbundling/</guid><guid>https://www.wired.com/story/prepare-for-doomsday/</guid><guid>https://www.wired.com/story/the-secret-essential-geography-of-the-office/</guid><guid>https://www.wired.com/story/usps-cdc-infrastructure-appreciation-society/</guid><guid>https://www.wired.com/story/course-catalog-school-year-2073-74/</guid><guid>https://www.wired.com/story/power-paradox-bad-software/</guid><guid>https://www.wired.com/story/databases-coding-real-programming-myth/</guid><guid>https://www.wired.com/story/infinite-loop-supply-chains/</guid><guid>https://www.wired.com/story/we-are-all-livestreamers-now-zoom-stage/</guid><guid>https://www.wired.com/story/stones-clocks-what-we-should-actually-leave-behind/</guid><guid>https://www.wired.com/story/millennials-genx-technology-explodes-generations/</guid><guid>https://www.wired.com/story/why-we-love-tech-defense-difficult-industry/</guid><guid>https://www.wired.com/2014/02/history/</guid>
```

`server` mode is enabled using the [aaronland/go-http-server](https://github.com/aaronland/go-http-server) package. Please consult that package's documentation for usage details.

#### Buckets

By default the `wired-feed` tool emits feeds to `STDOUT`.

It is also possible to output feed data to another source using the [GoCloud blob](https://gocloud.dev/howto/blob) abstraction layer by setting the `-bucket-uri` flag.

Supported blob targets are the local filesystem and AWS S3. Please consult that package's documentation for usage details.

## See also

* https://ftrain.com/
* https://github.com/gorilla/feeds
* https://gocloud.dev/howto/blob
* https://github.com/aaronland/go-http-server