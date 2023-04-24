package main

import "fmt"

func Message(name string) string {
	return "Hi, " + name
}

func main() {
	fmt.Println(Message("DUDE"))

}
