package main

import "fmt"

func main() {
	fmt.Printf("hello, world this is GO\n")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	age := 12

	if age < 13 {
		fmt.Println("Wow you're young!")
	}
	if age > 12 && age < 20 {
		fmt.Println("You're a teenager")
	}
	if age > 19 && age < 30 {
		fmt.Println("You're in your twenties")
	}
	if age > 29 && age < 40 {
		fmt.Println("you're in your thirties")
	}
	if age > 39 && age < 50 {
		fmt.Println("You're getting there!")
	}
	if age > 49 {
		fmt.Println("Over the hill!")
	}

}
