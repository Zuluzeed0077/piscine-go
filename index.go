package piscine

func Index(s, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i+len(toFind) <= len(s); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
