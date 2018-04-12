package main

import "fmt"

func main() {
	HelloWorld()
	ForLoop(10)
	ConditionalStatement(25)
	SwitchStatement(5)
	MaxPractice()
}

//HelloWorld prints a greeting
func HelloWorld() {
	fmt.Printf("hello, world this is GO\n")
}

//ForLoop prints numbers
func ForLoop(totalTimes int) {
	fmt.Println("for loop practice")
	for i := 0; i < totalTimes; i++ {
		fmt.Println(i)
	}
}

//ConditionalStatement given an age, prints a statement about that age
func ConditionalStatement(age int) {
	fmt.Println("conditional statement practice")

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

//SwitchStatement given a number representing a day, prints the associated value in english.
func SwitchStatement(day int) {

	fmt.Println("switch statement practice")
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

//MaxPractice calls the max function to get the max value of two integers and displays it
func MaxPractice() {
	fmt.Println("max practice")
	maxValue := Max(2, 7)
	fmt.Println(maxValue)
}

//Max determines the bigger value of the two and returns it
func Max(i int, j int) int {
	if i < j {
		return j
	}

	return i
}
