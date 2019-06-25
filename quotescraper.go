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

	fmt.Printf("\n");
	fmt.Printf("Filtered Results:");
	// Loop through to last page.
	// I have already determined the last page for this site.
	// TODO: detect 404 and stop loop when detected.
	i := 1
	for i < 11 {
		url := fmt.Sprintf("http://quotes.toscrape.com/page/%v", i);

		// See http://toscrape.com/ for more options!
		res, err := http.Get(url)
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

		// FIXME: Is this accumulative?
		filteredResults := doc.Find(selector).FilterFunction(func(i int, s *goquery.Selection) bool {
			return strings.Contains(s.Text(), filter);
		});

		filteredResults.Each(func(i int, s *goquery.Selection){
			fmt.Printf("\n");
			fmt.Printf("%v. %v\n", i+1, s.Text());
		});

		i += 1;
	}

}
