package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mmcdole/gofeed"
)

const baseUrl string = "http://heise.de/-"

type News struct {
	Date        string `json:"date,omitempty"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NewsList struct {
	News []News `json:"news,omitempty"`
	Url  string `json:"baseUrl,omitempty"`
}

func newsList(count int) NewsList {

	url := "https://www.heise.de/rss/heise-atom.xml"
	list := NewsList{
		Url: url,
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		fmt.Printf("error fetching the feed from url %s: %s\n", url, err)
		os.Exit(1)
	}
	for i := 0; i < count; i++ {
		item := feed.Items[i]
		id := strings.Replace(item.GUID, baseUrl, "", -1)
		news := News{
			Date:        formatDate(item.Published),
			ID:          id,
			Title:       item.Title,
			Description: item.Description,
		}
		list.News = append(list.News, news)
	}
	return list
}

func newsSingle(id string) (News, error) {

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")

	for _, item := range feed.Items {
		if item.GUID == "http://heise.de/-"+id {
			id := strings.Replace(item.GUID, baseUrl, "", -1)
			single := News{item.Published, id, item.Title, item.Description}
			return single, nil
		}
	}
	return News{}, errors.New("id not found")
}
