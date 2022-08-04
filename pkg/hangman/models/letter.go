package models

import "fmt"

type Letter struct {
	Value   rune
	Guessed bool
}

type Letters []Letter

func (letters Letters) Contains(char rune) bool {
	for _, l := range letters {
		if l.Value == char {
			return true
		}
	}
	return false
}

func (letters Letters) IsGuessed(char rune) bool {
	for _, l := range letters {
		if l.Value == char {
			return l.Guessed
		}
	}
	return false
}

func (letters Letters) Set(char rune, val bool) {
	for i, l := range letters {
		if l.Value == char {
			letters[i].Guessed = val
		}
	}
}

func (letters Letters) ToString() string {
	var out string
	for _, l := range letters {
		out += fmt.Sprintf("%c", l.Value)
	}

	return out
}

func (letters Letters) SetAll(val bool) {
	for i := range letters {
		letters[i].Guessed = val
	}
}
