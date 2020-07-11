package main

import "fmt"

func main() {

	/*
		//var card string = "Ace of Spades"
		card := "Ace of Spades" //:= should only be used while first initialisation only
		card = "Five of Diamonds"

	*/

	/*
			card := newCard()
			fmt.Println(card)

			cards := []string{newCard(), "Ace of Diamonds"}
			cards = append(cards, "Six of Spades")
			fmt.Println(cards)

			for i, card := range cards {
				fmt.Println(i, card)
			}


		cards := deck{newCard(), "Ace of Diamonds"}
		cards = append(cards, "Six of Spades")
		fmt.Println(cards)
	*/

	cards := newDeck()
	cards.print()

	hand, remainingDeck := deal(cards, 10)

	hand.print()
	remainingDeck.print()

	fmt.Println(cards.toString())

	cards.saveToFile("mycards")

	newCards := newDeckFromFile("mycards")
	newCards.shuffle()
	newCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
