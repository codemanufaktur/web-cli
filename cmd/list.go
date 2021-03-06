package cmd

import (
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

func list() {
	fp := gofeed.NewParser()
	url := "https://www.heise.de/rss/heise-atom.xml"
	feed, err := fp.ParseURL(url)
	if err != nil {
		fmt.Printf("error fetching the feed from url %s: %s\n", url, err)
		os.Exit(1)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(feed.Items[i].Title + " " + feed.Items[i].Description + " " + feed.Items[i].GUID)
	}
}
