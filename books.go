package main

import "fmt"

//Declare books Struct
type book struct {
	title string
	price int
}

//Create print method
func (b book) printPrice() {
	fmt.Println(b.price)
}
