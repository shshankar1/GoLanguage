package main

//var deckSize int = 20

//var money int

func main() {
	// var card string = "Ace of Spades"
	// var card = "Ace of Spades"
	//card := "Ace of Spades"
	//card = "Five of Diamonds"
	//card := newCard()
	//fmt.Println(card)
	//fmt.Println(deckSize)

	/*var deckNum int
	deckNum = 2
	fmt.Println(deckNum)

	money = 40
	fmt.Println(money)*/

	//fmt.Println(getTitle())

	//printState()

	//cards := []string{"Ace of Diamonds", newCard()}

	cards := deck{"Ace of Diamonds", newCard()}
	//fmt.Println(strings.Join(cards, ", "))

	cards.print()

	/*for idx := range cards {
		fmt.Println(idx)
	}*/
}

func newCard() string {
	return "Five of Diamonds"
}

func getTitle() string {
	return "Harry Potter"
}
