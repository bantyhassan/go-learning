package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Name: ")

	username, _ := reader.ReadString('\n')

	fmt.Println("Welcome ", username)

}
