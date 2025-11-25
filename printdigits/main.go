package main

import "github.com/01-edu/z01"

func main() {
	for z := '0'; z <= '9'; z++ {
		z01.PrintRune(z)
	}
	z01.PrintRune('\n')
}
