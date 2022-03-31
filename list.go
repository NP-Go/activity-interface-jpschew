package main

import "fmt"

// declare interface
type printer interface {
	printPrice()
}

// //declare List
type list []printer

// //declare method for List Print
func (l list) printPrice() {
	for _, item := range l {
		fmt.Println(item)
		item.printPrice()
	}

}
