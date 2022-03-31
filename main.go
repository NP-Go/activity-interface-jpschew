package main

func main() {
	//Declare objects
	game1 := game{"Minecraft", 5}
	game2 := game{"World of Warcraft", 19}
	game3 := game{"Elite Dangerours", 54}
	book1 := book{"Candle in the tomb", 20}
	book2 := book{"Barney and Friends", 10}
	compAcc1 := compAcc{"Razer BT earpiece", 159}
	compAcc2 := compAcc{"Razer keyboard", 110}
	compAcc3 := compAcc{"Logitech Mouse", 80}

	//Insert codes here
	var storeList list
	storeList = append(storeList, game1, game2, game3, book1, book2, compAcc1, compAcc2, compAcc3)
	storeList.printPrice()
}
