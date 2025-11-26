package piscine

func TrimAtoi(s string) int {
	n, sign := 0, 1
	digitStarted := false
	for _, r := range s {
		if r == '-' && !digitStarted {
			sign = -1
			continue
		}
		if r >= '0' && r <= '9' {
			digitStarted = true
			n = n*10 + int(r-'0')
		}
	}
	return n * sign
}
