package main

import (
	"fmt"
	poly "github.com/garrett-partenza-us/polynomial"
)

func main() {
	terms1 := []poly.Term{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	terms2 := []poly.Term{
		{3, 1},
		{4, 2},
		{5, 3},
		{6, -1},
	}

	poly1 := poly.Polynomial{Terms: terms1}
	poly2 := poly.Polynomial{Terms: terms2}

	fmt.Println(poly1.Add(poly2))
	fmt.Println(poly1.Sub(poly2))

	fmt.Println("Program ran successfully!")
}
