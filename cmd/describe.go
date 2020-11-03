package cmd

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"os"
	"time"
)

func describe(id string) {
	fp := gofeed.NewParser()
	url := "https://www.heise.de/rss/heise-atom.xml"
	feed, err := fp.ParseURL(url)
	if err != nil {
		fmt.Printf("error fetching the feed from url %s: %s\n", url, err)
		os.Exit(1)
	}
	for _, element := range feed.Items {
		if element.GUID == id {
			fmt.Print(element.Title)
			fmt.Print(" ## ")
			fmt.Print(element.Description)
			fmt.Print(" ## ")
			fmt.Println(formatDate(element.Published))
			return
		}
	}
}

func formatDate(date string) string {
	ref := "2006-01-02T15:04:05-07:00"
	t, err := time.Parse(ref, date)
	if err != nil {
		return ""
	}
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		fmt.Printf("error loading location 'Europe/Berlin': %s\n", err)
		os.Exit(1)
	}
	return t.In(loc).Format("02.01.2006 um 15:04")
}
