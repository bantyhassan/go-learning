package main

import "fmt"

type address struct {
	street     string
	city       string
	postalCode string
}

type emplyoee struct {
	firstName string
	lastName  string
	address   address
}

func main() {

	emp1 := emplyoee{"Hassan", "Ahamed", address{"Lalkuthi Road", "Midnapore", "721101"}}

	emp1.print()

}

func (e emplyoee) print() {
	// fmt.Println(e)
	fmt.Printf("%+v", e)
}
