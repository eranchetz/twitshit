package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	// Generator will create a new instance with an array of startup name and an array with ~5000 nouns
	gen, err := New()
	if err != nil {
		log.Fatalln("could not create generator", err)
	}

	fmt.Println(gen.StartUps)
	fmt.Println("********************************")
	fmt.Println(gen.Words)

	// Build tweet
	tweetBody := fmt.Sprintln(gen.StartUps[rand.Intn(len(gen.StartUps))], "for", gen.Words[rand.Intn(len(gen.Words))])

	// Tweet shit it
	gen.Tweet(tweetBody)

}
