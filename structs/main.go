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

	fmt.Println(emp1)

}
