package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Generator struct {
	StartUps        []string
	Words           []string
	consumerKey     string
	consumerSecret  string
	accessKey       string
	accessKeySecret string
}

func init() {
	rand.Seed(time.Now().Unix())
}

func New() (*Generator, error) {
	gen := new(Generator)
	doc, err := goquery.NewDocument("http://www.startupranking.com/top/united-states")
	if err != nil {
		return nil, err
	}

	gen.consumerKey = os.Getenv("CONSUMER_KEY")
	gen.consumerSecret = os.Getenv("CONSUMER_SECRET")
	gen.accessKey = os.Getenv("ACCESS_TOKEN")
	gen.accessKeySecret = os.Getenv("ACCESS_TOKEN_SECRET")

	// Find the startup names
	doc.Find(".name").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the name
		startup := s.Text()
		//fmt.Printf("Name %d - %s\n", i, startup)
		gen.StartUps = append(gen.StartUps, startup)
	})

	// Load all uncountabl words to memory
	gen.Words, err = readLines(`./wordlist/unc.txt`)
	if err != nil {
		return nil, err
	}

	return gen, nil
}

func (g *Generator) Tweet(tweet string) {
	fmt.Println(g.accessKey, g.accessKeySecret, g.consumerKey, g.consumerSecret)
	config := oauth1.NewConfig(g.consumerKey, g.consumerSecret)
	token := oauth1.NewToken(g.accessKey, g.accessKeySecret)
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	_, resp, err := client.Statuses.Update(tweet, nil)
	if err != nil {
		fmt.Println("twitter update error", err)
	}
	fmt.Println(resp)

}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
