package main

import "fmt"

//Declare Computer Accessoies Struct
type compAcc struct {
	title string
	price int
}

//Create print method
func (ca compAcc) printPrice() {
	fmt.Println(ca.price)
}
