package main

import (
	"fmt"
)

const (
	msgPrefixEnglish = "Hi, "
	msgPrefixSpanish = "Hola, "
	msgPrefixFrench  = "Bonjour, "
)

const (
	spanish = "Spanish"
	french  = "French"
	english = "English"
)

func Message(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return switchPrefix(lang) + name
}

func switchPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = msgPrefixSpanish
	case french:
		prefix = msgPrefixFrench
	default:
		prefix = msgPrefixEnglish
	}
	return prefix
}

func main() {
	fmt.Println(Message("DUDE", "Spanish"))
}
