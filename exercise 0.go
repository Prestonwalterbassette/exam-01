package main

import "fmt"

func main() {
	for i := 'a' ; i <= 'z'; i++ {
		fmt.Fprint(i)
	}
	fmt.Fprint(\n)
}