package main

import "fmt"

func main() {
	input := "MCMXCIV"
	output := romanToIntSolution2(input)
	fmt.Println(output)
}

func romanToIntSolution1(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0

	previousChar := ""
	for i := 0; i < len(s); i++ {
		char := string([]rune(s)[i])
		if previousChar == "" {
			previousChar = char
			continue
		} else if romanMap[char] > romanMap[previousChar] {
			result += romanMap[char] - romanMap[previousChar]
			previousChar = ""
			continue
		} else {
			result += romanMap[previousChar]
			previousChar = char
		}
	}

	if previousChar != "" {
		result += romanMap[previousChar]
		previousChar = ""
	}

	return result
}

func romanToIntSolution2(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	i := 0
	for i < len(s) {
		// subtractive case
		if i+1 < len(s) && romanMap[rune(s[i])] < romanMap[rune(s[i+1])] {
			result += romanMap[rune(s[i+1])] - romanMap[rune(s[i])]
			i += 2
		} else {
			result += romanMap[rune(s[i])]
			i += 1
		}
	}

	return result
}
