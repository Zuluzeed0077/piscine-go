package main

func IsAlpha(s string) bool {
	for _, r := range s {
		if !((r >= 'A' && r <= 'Z') ||
			(r >= 'a' && r <= 'z') ||
			(r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}
func main() {
	s := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}
