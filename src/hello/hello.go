package main

import "fmt"

func main() {
	fmt.Printf("hello, world this is GO\n")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	age := 39

	if age < 13 {
		fmt.Println("Wow you're young!")
	} else if age < 20 {
		fmt.Println("You're a teenager")
	} else if age < 30 {
		fmt.Println("You're in your twenties")
	} else if age < 40 {
		fmt.Println("you're in your thirties")
	} else if age < 50 {
		fmt.Println("You're getting there!")
	} else {
		fmt.Println("Over the hill!")
	}

}
