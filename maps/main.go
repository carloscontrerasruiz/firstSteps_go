package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "123456",
		"blue": "67890",
	}

	//colors := make(map[string]string)

	//var colors map[string]string

	colors["white"] = "123456"

	//delete(colors, "blue")
	printMap(colors)
	//fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("HEx code for ", color, " is ", hex)
	}
}
