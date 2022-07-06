package main

import "fmt"

func romanToInt(s string) int {
	map1 := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	num := map1[string(s[0])]
	l := len(s)
	for i := 1; i < l; i++ {
		cx := string(s[i])
		add := map1[cx]
		switch cx {
		case "V", "X":

			if s[i-1] == 'I' {
				add -= 2
			}
		case "L", "C":

			if s[i-1] == 'X' {
				add -= 20
			}
		case "D", "M":

			if s[i-1] == 'C' {
				add -= 200
			}
		}
		num += add
	}
	return num
}

func main() {
	fmt.Println(romanToInt("MCXII"))

}
