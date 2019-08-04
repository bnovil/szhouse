package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	log.Print("===== start =====")
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	t, _ := doc.Find(".table-wrap").Find(".ms-link").Attr("href")
	bs, _ := json.Marshal(t)
	fmt.Println(string(bs))

	s := doc.Find(".sidebar-reviews").Find(".header-sm").Text
	fmt.Printf("%s\n", s)

	s2 := doc.Find(".sidebar-reviews").Find("a").Text
	fmt.Println(s2)

	// Find the review items
	doc.Find(".sidebar-reviews .clearfix .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find(".header-xs").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})



}

func main() {
	ExampleScrape()
}
