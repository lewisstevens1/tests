package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

var fileStringDelimiter = ","

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}

	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	// Return :handSize (0 to Handsize)
	// Return handSize: (Handsize to Deck length)
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), fileStringDelimiter)
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), fileStringDelimiter)
	return deck(s)
}

func (d deck) shuffle() {
	time := time.Now().UnixNano()
	source := rand.NewSource(time)
	random := rand.New(source)

	for i := range d {
		newPosition := random.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
