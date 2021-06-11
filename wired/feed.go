package wired

import (
	"context"
	"fmt"
	"github.com/gorilla/feeds"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"time"
)

// GenerateFeed will create a mmcdole/gofeed/rss.Feed instance of Paul Ford's essays on the WIRED website. The number of items will be capped at 'max_items'. If 'max_item' is -1 then all the essays will be included.
func GenerateFeed(ctx context.Context, max_items int) (*feeds.Feed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", URL_ESSAYS, nil)

	if err != nil {
		return nil, fmt.Errorf("Unable to generate feed, %w", err)
	}

	cl := &http.Client{}
	rsp, err := cl.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Failed to request latest essays page, %w", err)
	}

	defer rsp.Body.Close()

	return GenerateFeedWithReader(ctx, rsp.Body, max_items)
}

// GenerateFeed will create a mmcdole/gofeed/rss.Feed instance of Paul Ford's essays included in the HTML document in 'r'. The number of items will be capped at 'max_items'. If 'max_item' is -1 then all the essays will be included.
func GenerateFeedWithReader(ctx context.Context, r io.Reader, max_items int) (*feeds.Feed, error) {

	doc, err := html.Parse(r)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse HTML, %w", err)
	}

	items := make([]*feeds.Item, 0)

	is_link := false
	title := ""
	link := ""
	date := ""
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

			if n.Type == html.ElementNode && n.Data == "div" {

				for _, a := range n.Attr {

					switch a.Key {
					case "class":

						if a.Val == "summary-item__publish-date" {
							date = n.FirstChild.Data
							break
						}
					default:
						// pass
					}
				}

			}

		}

		if is_link && title != "" && link != "" && desc != "" && date != "" {

			item_link := &feeds.Link{
				Href: URL_WIRED + link,
			}

			item := &feeds.Item{
				Title:       title,
				Link:        item_link,
				Description: desc,
				Id:          URL_WIRED + link,
			}

			t, err := time.Parse("01.02.2006 15:04 AM", date)

			if err != nil {
				log.Printf("Failed to parse date '%s' for %s, %v", date, link, err)
			} else {
				item.Created = t
			}

			items = append(items, item)

			is_link = false
			title = ""
			link = ""
			desc = ""
			date = ""
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

	feed_link := &feeds.Link{
		Href: "https://www.wired.com/author/paul-ford/",
	}

	now := time.Now()

	feed := &feeds.Feed{
		Title:       "Paul Ford's WIRED essays",
		Description: "Paul Ford's WIRED essays",
		Link:        feed_link,
		Items:       items,
		Created:     now,
	}

	return feed, nil
}
