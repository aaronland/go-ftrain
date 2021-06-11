# go-ftrain

Go package for working with things related to Paul Ford.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/go-ftrain.svg)](https://pkg.go.dev/github.com/aaronland/go-ftrain)

## Tools

### wired-rss-feed

Generate an RSS feed of Paul Ford's WIRED essays.

```
$> go run -mod vendor cmd/wired-rss-feed/main.go | jq '.items[]["title"]'

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