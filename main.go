package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generator will create a new instance with an array of startup name and an array with ~5000 nouns
	gen, _ := New()

	// Build tweet
	tweetBody := fmt.Sprintln(gen.startUps[rand.Intn(len(gen.startUps))], "for", gen.words[rand.Intn(len(gen.words))])

	// Tweet shit it
	gen.Tweet(tweetBody)

}
