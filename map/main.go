package main

import "fmt"

type colorMapType map[string]string

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"
	// colors["black"] = "#000000"
	// delete(colors, "white")

	colors := colorMapType{
		"red":   "ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//map type is a reference type, so it does not needs pointers to mutate
	colors.addColor("black", "#000000")
	printMap(colors)

	//Only recap with pointers
	//string type is a value type, so it needs pointers to mutate
	test := "Teststring"
	changeText(&test)
	fmt.Println(test)

}

func (m colorMapType) addColor(color string, hex string) {
	m[color] = hex
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v color has hex number %v\n", color, hex)
	}
}

func changeText(s *string) {
	(*s) = (*s) + "!!!"
}
