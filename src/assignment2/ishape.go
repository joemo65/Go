package main

import (
	"fmt"
)

type ishape interface {
	getArea() float64
}

func printArea(shape ishape) {
	fmt.Println(shape.getArea())
}
