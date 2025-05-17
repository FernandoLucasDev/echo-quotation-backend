package services

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"api-echo/db"
	"api-echo/models"

	"gorm.io/gorm"
)

type RSS struct {
	Channel struct {
		Items []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Enclosure   struct {
		URL string `xml:"url,attr"`
	} `xml:"enclosure"`
	MediaContent struct {
		URL string `xml:"url,attr"`
	} `xml:"http://search.yahoo.com/mrss/ content"`
}


var rssFeeds = map[string]string{
	"G1":           "https://g1.globo.com/rss/g1/economia/",
	"Gazeta do Povo": "https://www.gazetadopovo.com.br/feed/rss/economia.xml",
	"BBC":          "https://feeds.bbci.co.uk/news/world/rss.xml",
}

func FetchAndStoreNews() {
	for midia, url := range rssFeeds {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Erro ao acessar:", url, err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erro ao ler conteúdo de:", url, err)
			continue
		}

		var rss RSS
		if err := xml.Unmarshal(body, &rss); err != nil {
			fmt.Println("Erro ao parsear XML:", err)
			continue
		}

		for _, item := range rss.Channel.Items {
			news := models.News{
				Title: item.Title,
				Link:  item.Link,
				Midia: midia,
				Image: item.MediaContent.URL,
				CreatedAt: time.Now(),
			}

			exists := models.News{}
			result := db.DB.Where("link = ?", news.Link).First(&exists)
			if result.Error == gorm.ErrRecordNotFound {
				db.DB.Create(&news)
			}
		}
	}
	threshold := time.Now().AddDate(0, -1, 0)
	db.DB.Where("created_at < ?", threshold).Delete(&models.News{})

	fmt.Println("Job de notícias executado em", time.Now())
}
