package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func scraper() {
	// Make HTTP GET request
	response, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal("Error fetching Google webpage:", err)
	}
	defer response.Body.Close()

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error parsing Google webpage:", err)
	}

	// Find and print the title
	title := doc.Find("title").Text()
	fmt.Printf("Hi, the title element of Google's webpage is: %s\n", title)
}

func main() {
	scraper()
	fmt.Println("Press Enter to close...")
	fmt.Scanln() // Wait for Enter key press
}
