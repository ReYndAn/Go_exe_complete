package main

import "fmt"

func romanToArabic(romanNum string) int {
	romanToArabicMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	arabicNum := 0
	for i := 0; i < len(romanNum); i++ {
		if i+1 < len(romanNum) && romanToArabicMap[romanNum[i:i+2]] > 0 {
			arabicNum += romanToArabicMap[romanNum[i:i+2]]
			i++
		} else {
			arabicNum += romanToArabicMap[string(romanNum[i])]
		}
	}
	return arabicNum
}

func main() {
	romanNum := "XXX"
	arabicNum := romanToArabic(romanNum)
	fmt.Printf("%s in roman numerals is %d in arabic numerals\n", romanNum, arabicNum)
}
