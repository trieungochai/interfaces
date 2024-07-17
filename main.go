package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "hi there!"
}

func (spanishBot) getGreeting() string {
	return "hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
