package main

import "fmt"

func main() {
	a := 24
	b := 36

	for b != 0 {
		a, b = b, a%b

	}
	fmt.Print("Le PGCD est ", a)
}