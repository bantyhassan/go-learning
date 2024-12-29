package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[string]string)

	colors["red"] = "#FF0000"

	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"green": "#3da26c",
	// }

	fmt.Println(colors)
}
