package main

import (
	"fmt"
)

const (
	msgPrefixEnglish = "Hi, "
	msgPrefixSpanish = "Hola, "
	msgPrefixFrench  = "Bonjour, "
)

// const (
// 	english = "English"
// 	spanish = "Spanish"
// 	french  = "French"
// )

func Message(name, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == "Spanish" {
		return msgPrefixSpanish + name
	}
	return msgPrefixEnglish + name
}

func main() {
	fmt.Println(Message("DUDE", "Spanish"))
}
