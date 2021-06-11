package wired

import (
	"context"
	"io"
	"github.com/mmcdole/gofeed/rss"
	"golang.org/x/net/html"
	"fmt"
	"net/http"
	_ "log"
)

const URL_WIRED string = "https://www.wired.com"

const URL_ESSAYS string = "https://www.wired.com/author/paul-ford/"

func GenerateRSSFeed(ctx context.Context, max_items int) (*rss.Feed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", URL_ESSAYS, nil)

	if err != nil {
		return nil, fmt.Errorf("Unable to generate RSS feed, %w", err)
	}

	cl := &http.Client{}
	rsp, err := cl.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Failed to request latest essays page, %w", err)
	}

	defer rsp.Body.Close()

	return GenerateRSSFeedWithReader(ctx, rsp.Body, max_items)
}

func GenerateRSSFeedWithReader(ctx context.Context, r io.Reader, max_items int) (*rss.Feed, error) {

	doc, err := html.Parse(r)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse HTML, %w", err)
	}

	items := make([]*rss.Item, 0)

	is_link := false
	title := ""	
	link := ""
	desc := ""
	
	var f func(*html.Node)

	f = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {

			for _, a := range n.Attr {

				switch a.Key {
				case "class":
					
					if a.Val == "summary-item-tracking__hed-link summary-item__hed-link" {
						is_link = true
					}
					
				case "href":
					link = a.Val
				default:
					// pass
				}
			}
			
		}

		if is_link {

			if n.Type == html.ElementNode && n.Data == "h2" {
				title = n.FirstChild.Data
			}

			if n.Type == html.ElementNode && n.Data == "p" {
				desc = n.FirstChild.Data
			}
		}

		if is_link && title != "" && link != "" && desc != "" {

			item := &rss.Item{
				Title: title,
				Link: URL_WIRED + link,
				Description: desc,
			}

			items = append(items, item)

			is_link = false
			title = ""
			link = ""
			desc = ""
		}

		if max_items > 0 && len(items) == max_items {
			return
		}
		
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}		
	}

	f(doc)

	if len(items) == 0 {
		return nil, fmt.Errorf("Unable to derive any feed items")
	}

	feed := &rss.Feed{
		Title: "Paul Ford's WIRED essays",
		Link: "https://www.wired.com/author/paul-ford/",
		Version: "2.0",
		Items: items,
	}
	
	return feed, nil
}
