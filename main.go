package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	scrapeCourses("http://www.calendar.ubc.ca/vancouver/courses.cfm?page=name")
}

type course struct {
	SIS  string
	Name string
}

func scrapeCourses(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalln("could not build doc: ", err)
	}

	doc.Find("#UbcMainContent .row-highlight").Children().Children().Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists || !strings.HasPrefix(href, "courses.cfm") {
			return
		}
		fmt.Println(s.Text())
	})
}
