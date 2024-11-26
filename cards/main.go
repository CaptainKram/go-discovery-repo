package main

func main() {
	cards := newDeck()

	hand1, cards := deal(cards, 5)

	hand1.print()
	cards.print()
}
