package main

// import "io"
import "fmt"

func main() {

	// // 1. A way to assign a variable
	// var card string = "Ace of Spades"

	// // 2. A way to assign a variable
	// card := newCard()

	// // 3. A way to override previous variable
	// card = "Five of diamonds"

	// // 4. A way to append to a list
	// cards := []string{newCard(), newCard()}
	// cards = append(cards, "Six of Spades")

	// // 5. A way to loop over a list
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// // 6. A way to call a function
	// cards := newDeck()

	// // 7. A way to call a function and return two seperate values
	// hand, remainingCards := deal(cards, 5)

	// // 8. String conversion so that it can be written to a file
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// newCards := newDeckFromFile("my_cards")
	// fmt.Println("New Cards:", newCards)

	// newCardsError := cards.newDeckFromFile("my_cards_error")
	// fmt.Println("New Cards Error:", newCardsError)

	// 9. Shuffling cards
	cards := newDeck()
	fmt.Println(cards)
	cards.shuffle()
	fmt.Println(cards)
}

// Passing as receiver vs argument
// e.g. dave.funcA() vs funcA(dave)
//
