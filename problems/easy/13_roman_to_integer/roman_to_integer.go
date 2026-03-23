package romantointeger

func romanToInt(s string) int {
	chars := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && chars[s[i]] < chars[s[i+1]] {
			sum -= chars[s[i]]
		} else {
			sum += chars[s[i]]
		}
	}
	return sum
}
