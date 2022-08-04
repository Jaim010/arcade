package file

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func GetRandomWord() string {
	words := getWords()

	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(len(words))
	return words[randInt]
}

func getWords() []string {
	var words []string

	file, err := os.Open("../config/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}
