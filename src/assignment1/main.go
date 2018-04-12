package main

import "fmt"

func main() {
	s := createSlice(10)

	printSlice(s)
}

func createSlice(length int) []int {
	s := []int{}

	for i := 0; i < length; i++ {
		s = append(s, i)
	}

	return s
}

func printSlice(slice []int) {
	for i := range slice {
		if slice[i]%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

}
