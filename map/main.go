package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b1726",
		"white": "#ffffff",
	}

	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	// fmt.Println(colors)

	printMap(colors)
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " code is " + hex)
		c[color] = color + "x"
	}
}
