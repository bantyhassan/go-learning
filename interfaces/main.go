package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func (eb englishBot) getGreeting() string {
// 	// Very custom logic for generating greeting in english
// 	return "Hello World!"
// }

// func (sb spanishBot) getGreeting() string {
// 	// Very custom logic for generating greeting in spanish
// 	return "��Hola Mundo!"
// }

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// Very custom logic for generating greeting in english
	return "Hello World!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating greeting in spanish
	return "��Hola Mundo!"
}
