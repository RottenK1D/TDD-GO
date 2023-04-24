package main

import "fmt"

const msgPrefix = "Hi, "

func Message(name string) string {
	if name == "" {
		name = "World"
	}
	return msgPrefix + name
}

func main() {
	fmt.Println(Message("DUDE"))
}
