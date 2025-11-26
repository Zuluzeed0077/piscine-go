package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		seen[r] = true
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	bLen := len(base)
	var digits []rune
	for nbr > 0 {
		rem := nbr % bLen
		digits = append(digits, rune(base[rem]))
		nbr = nbr / bLen
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
