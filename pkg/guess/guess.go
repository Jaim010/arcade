package guess

import (
	"fmt"
	"strings"
)

type GuessKind int64

const (
	Letter GuessKind = iota
	Word
)

type Guess struct {
	Value string
	Kind  GuessKind
}

func GetGuess() (Guess, error) {
	fmt.Print("Guess a letter / Guess the word: ")
	var guess string
	fmt.Scanln(&guess)

	var kind GuessKind
	switch guessLen := len(guess); {
	case guessLen <= 0:
		return Guess{}, fmt.Errorf("Empty value")
	case guessLen == 1:
		kind = Letter
	default:
		kind = Word
	}

	guess = strings.ToLower(guess)
	return Guess{Kind: kind, Value: guess}, nil
}

func (kind GuessKind) equals(val GuessKind) bool {
	return kind == val
}

func (kind GuessKind) IsLetter() bool {
	return kind.equals(Letter)
}

func (kind GuessKind) IsWord() bool {
	return kind.equals(Word)
}
