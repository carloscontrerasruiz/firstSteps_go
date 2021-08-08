package main

func main() {
	//var card string = "Ace of Spades"
	//card2 := newCard()
	//card2 = "Ace of Spades 2"
	//cards := [2]string{newCard(), newCard()} //array

	//cards := []string{newCard(), newCard()} //slice
	cards := newDeck()

	//hand1, hand2 := deal(cards, 5)

	//hand1.print()
	//hand2.print()

	//cards.print()
	//cards.saveToFile("card")

	//cards2 := newDeckFromFile("cards")

	//cards2.print()
	//fmt.Println(cards.toString())

	cards.shufleDeck()
	cards.print()
}
