package main

import (
	"fmt"
)

func main() {
	colors := make(map[string]string)
	colors["white"] = "#ffffff"

	//delete(colors, "white")
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	fmt.Println(colors)
}
