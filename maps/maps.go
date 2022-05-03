package main

import "fmt"

func main() {
	// Maps always use square brace syntax due to keys being anytime e.g. 10 could be a key name.

	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"green": "#00FF00",
	}

	// // // This:
	// // var colors map[string]string
	// // // is equal to:
	// colors := make(map[string]string)
	// colors["white"] = "#FFFFFF"

	printMap(colors)
	// fmt.Println(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color, hex)
	}
}
