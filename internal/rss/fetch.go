package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		feedURL,
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rssBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	if err := xml.Unmarshal(rssBytes, &rssFeed); err != nil {
		return nil, err
	}

	for idx := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[idx].Title = html.UnescapeString(rssFeed.Channel.Item[idx].Title)
		rssFeed.Channel.Item[idx].Description = html.UnescapeString(rssFeed.Channel.Item[idx].Description)
	}

	return &rssFeed, nil
}
