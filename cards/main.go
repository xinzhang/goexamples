package main

func main() {
	cards := newDeck()
	cards = cards.NewCard("joker")
	//cards = append(cards, "six of heart")
	cards.PrintCard()
}
