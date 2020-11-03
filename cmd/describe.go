package cmd

import (
	"fmt"
	"os"
	"time"
)

func describe(id string) {

	news, err := newsSingle(id)
	if err != nil {
		fmt.Printf("error fetching the feed: %s", err)
		os.Exit(1)
	}

	fmt.Print(news.Title)
	fmt.Print(" ## ")
	fmt.Print(news.Description)
	fmt.Print(" ## ")
	fmt.Print(formatDate(news.Date))
	fmt.Print(" ## ")
	fmt.Print(news.ID)
}

func describeHtmlTable(id string) string {
	news, err := newsSingle(id)
	if err != nil {
		fmt.Printf("error fetching the feed: %s", err)
		os.Exit(1)
	}

	item := fmt.Sprintf("<table><tr><th>%s</th><th>%s</th><th>%s</th><th>%s</th></tr></table>", news.Title, news.Description, formatDate(news.Date), news.ID)
	return item
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
