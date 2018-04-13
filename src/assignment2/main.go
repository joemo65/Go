package main

func main() {
	t := triangle{
		base:   2,
		height: 5,
	}

	s := square{
		sideLength: 5,
	}

	printArea(t)
	printArea(s)
}
