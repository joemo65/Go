package main

import "fmt"

func main() {
	fmt.Printf("hello, world this is GO\n")

	fmt.Println("for loop practice")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("conditional statement practice")
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

	fmt.Println("switch statement practice")
	day := 0
	switch day {
	case 0:
		fmt.Println("sunday")
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	default:
		fmt.Println("it's groundhog day")
	}
}
