package game

import (
	"fmt"
	"strings"

	"github.com/Jaim010/hangman/pkg/file"
	"github.com/Jaim010/hangman/pkg/guess"
	"github.com/Jaim010/hangman/pkg/models"
	"github.com/Jaim010/hangman/pkg/ui"
	"github.com/fatih/color"
)

const letters = "qwertyuiopasdfghjklzxcvbnm"

var (
	green  = color.New(color.FgGreen).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	black  = color.New(color.FgBlack).SprintFunc()
)

type GameMode int64

const (
	Easy GameMode = iota
	Hard
)

type game struct {
	mode     GameMode
	letters  models.Letters
	word     models.Letters
	mistakes uint8
	hangman  ui.Hangman
	win      bool
	running  bool
	msg      string
}

func New() game {
	return game{}
}

func (game *game) Run() {
	for {
		game.setup()
		for game.running {
			ui.Clear()
			game.hangman.Print()
			game.printWord()
			game.printLetters()
			if game.msg != "" {
				fmt.Println(game.msg)
				game.msg = ""
			}
			guess, err := guess.GetGuess()
			if err != nil {
				continue
			}

			game.handle(guess)
			game.hangman.Update(game.mistakes)
			game.check()
		}

		// Last UI update
		ui.Clear()
		game.hangman.Print()
		game.printWord()
		game.printLetters()

		fmt.Printf("The word was: '%s'\n", yellow(game.word.ToString()))
		if game.win {
			fmt.Printf("You %s!\n", green("won"))
		} else {
			fmt.Printf("You %s!\n", red("lost"))
		}

		if !game.Restart() {
			break
		}
	}
}

func (game *game) setup() {
	game.running = true
	game.win = false
	game.msg = ""
	game.mistakes = 0

	// Get random word
	game.word = []models.Letter{}

	// Setup word
	word := file.GetRandomWord()
	word = strings.ToLower(word)
	for _, char := range word {
		game.word = append(game.word, models.Letter{Value: char, Guessed: false})
	}

	game.hangman = ui.Hangman{}

	// Setup guessed letters
	game.letters = []models.Letter{}
	for _, char := range letters {
		game.letters = append(game.letters, models.Letter{Value: char, Guessed: false})
	}
}

func (game *game) check() {
	if game.mistakes >= 10 {
		game.running = false
		return
	}

	for _, l := range game.word {
		if !l.Guessed {
			game.running = true
			return
		}
	}

	game.win = true
	game.running = false
}

func (game game) printWord() {
	out := ""
	for _, letter := range game.word {
		if letter.Guessed {
			out += fmt.Sprintf("%c ", letter.Value)
		} else {
			out += "_ "
		}
	}
	fmt.Println(out + "\n")
}

func (game game) printLetters() {

	var out string
	for _, l := range game.letters {

		char := fmt.Sprintf("%c", l.Value)
		if l.Guessed {
			out += black(char) + " "
		} else {
			out += green(char) + " "
		}

		switch l.Value {
		case 'p':
			out += "\n "
		case 'l':
			out += "\n  "
		case 'm':
			out += "\n"
		}
	}

	fmt.Println(out)
}

func (game game) Restart() bool {
	var stop bool
	for {
		var input string
		fmt.Printf("Play again? %s/%s ", green("y"), red("n"))
		fmt.Scanln(&input)

		if len(input) == 1 {
			if input == "y" {
				stop = true
				break
			} else if input == "n" {
				stop = false
				break
			}
		}
	}

	return stop
}
