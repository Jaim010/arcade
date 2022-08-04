package game

import (
	"fmt"

	"github.com/Jaim010/arcade/pkg/hangman/guess"
)

func (game *game) handle(guess guess.Guess) {
	if guess.Kind.IsLetter() {
		// Since it's a enum with type Letter
		// We can assume it's an string with length 1
		letter := []rune(guess.Value)[0]
		game.handleLetter(letter)
	} else if guess.Kind.IsWord() {
		game.handleWord(guess.Value)
	}
}

func (game *game) handleLetter(letter rune) {
	if !game.letters.IsGuessed(letter) {
		game.letters.Set(letter, true)

		var result string
		if game.word.Contains(letter) {
			game.word.Set(letter, true)
			result = green("correct")
		} else {
			game.mistakes += 1
			result = red("wrong")
		}
		charStr := fmt.Sprintf("%c", letter)
		game.msg = fmt.Sprintf("'%s' was %s!", yellow(charStr), result)
	} else {
		charStr := fmt.Sprintf("%c", letter)
		game.msg = fmt.Sprintf("'%s' was %s!", yellow(charStr), red("already guessed"))
	}
}

func (game *game) handleWord(word string) {
	game.win = word == game.word.ToString()
	if !game.win {
		game.mistakes += 1
	}

	if game.win || game.mode == Hard {
		if game.win {
			// Sets all the guessed letters to true
			// Required for win condition
			game.word.SetAll(true)
			for _, l := range game.word {
				game.letters.Set(l.Value, true)
			}
		}

		game.running = false
		return
	}

	game.msg = fmt.Sprintf("'%s' was %s", yellow(word), red("wrong"))
}
