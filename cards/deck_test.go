package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	dl := 16

	if len(d) != dl {
		t.Errorf("Expected deck length of %d, but got %d", dl, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	dl := 16

	d := newDeck()

	os.Remove(filename)
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != dl {
		t.Errorf("Expected deck length of %d, but got %d", dl, len(loadedDeck))
	}

	os.Remove(filename)
}
