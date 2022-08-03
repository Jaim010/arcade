package ui

import "fmt"

type Hangman struct {
	Visual [7]string
}

func (hangman *Hangman) Update(mistakes uint8) Hangman {
	switch mistakes {
	case 1:
		hangman.Visual[6] = "__________"
	case 2:
		hangman.Visual[6] = "____||____"
		for i := 1; i < len(hangman.Visual)-1; i++ {
			hangman.Visual[i] = "    ||    "
		}
	case 3:
		hangman.Visual[0] = "   ____________"
	case 4:
		hangman.Visual[1] = "    ||     |   "
	case 5:
		hangman.Visual[2] = "    ||     O"
	case 6:
		hangman.Visual[3] = "    ||     |"
		hangman.Visual[4] = "    ||     |"
	case 7:
		hangman.Visual[3] = "    ||    \\|"
	case 8:
		hangman.Visual[3] = "    ||    \\|/"
	case 9:
		hangman.Visual[5] = "    ||    /"
	case 10:
		hangman.Visual[5] = "    ||    / \\"
	}

	return *hangman
}

func (hangman Hangman) Print() {
	for _, row := range hangman.Visual {
		fmt.Println(row)
	}
	fmt.Println()
}
