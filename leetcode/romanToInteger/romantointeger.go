package romanToInteger

func RomanToInteger(s string) int {
	result := 0

	var romanIntegers = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) &&
			romanIntegers[string(s[i])] < romanIntegers[string(s[i+1])] {
			result -= romanIntegers[string(s[i])]
		} else {
			result += romanIntegers[string(s[i])]
		}
	}

	return result
}
