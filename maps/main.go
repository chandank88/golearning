package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}

	printMap(colors)
	//var colors map[string]string
	/*colors := make(map[string]string)

	colors["white"] = "#ffffff"

	delete(colors, "white")*/
	fmt.Println(colors)
}
