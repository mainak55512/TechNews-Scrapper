package scrapper

import (
	"github.com/gocolly/colly"
)

func NewsScrapper() [][][]string {
	var records [][][]string
	records = append(records, wiredScrapper(), infoworldScrapper(),developertechScrapper())
	return records
}

func wiredScrapper() [][]string {
	var records [][]string
	wired_url := "https://www.wired.com/tag/programming/"
	c := colly.NewCollector()
	c.OnHTML(".kRJmrk", func(e *colly.HTMLElement) {
		records = append(records, []string{
			e.ChildText(".fcczqx"),
			"https://www.wired.com" + e.Attr("href"),
		})
	})
	c.Visit(wired_url)
	return records
}

func infoworldScrapper() [][]string {
	var records [][]string
	infoworld_url := "https://www.infoworld.com/category/programming-languages/"
	c := colly.NewCollector()
	c.OnHTML(".river-well", func(e *colly.HTMLElement) {
		records = append(records, []string{
			e.ChildText(".post-cont h3"),
			"https://www.infoworld.com" + e.ChildAttr(".post-cont h3 a", "href"),
		})
	})
	c.Visit(infoworld_url)
	return records
}

func developertechScrapper() [][]string{
	var records [][]string
	infoworld_url := "https://www.developer-tech.com/"
	c := colly.NewCollector()
	c.OnHTML(".article-header", func(e *colly.HTMLElement) {
		records = append(records, []string{
			e.ChildText("h3"),
			e.ChildAttr("h3 a", "href"),
		})
	})
	c.Visit(infoworld_url)
	return records
}
