package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"bufio"
	"os"
);

func main () {

	// Get input to filter results
	fmt.Printf("Search Quotes: ");
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filter := scanner.Text();
	// See http://toscrape.com/ for more options!
	res, err := http.Get("http://quotes.toscrape.com/")
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

	selector := "[itemprop='text']"

	filteredResults := doc.Find(selector).FilterFunction(func(i int, s *goquery.Selection) bool {
		return strings.Contains(s.Text(), filter);
	});

	fmt.Printf("\n");
	fmt.Printf("Filtered Results:");

	filteredResults.Each(func(i int, s *goquery.Selection){
		fmt.Printf("\n");
		fmt.Printf("%v. %v\n", i+1, s.Text());
	});
}
