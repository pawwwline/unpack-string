package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := unpackString("")
	fmt.Println(str)

}
func unpackString(str string) string {
	res := make([]byte, 0)
	ints := map[rune]struct{}{'1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {}, '0': {}}
	l, r := 0, 0
	for r < len(str) {
		_, isDigit := ints[rune(str[r])]
		if !isDigit {
			l = r
			res = append(res, str[l])
			r++
			continue
		}
		digitStart := r
		for r < len(str) {
			_, isDigit = ints[rune(str[r])]
			if !isDigit {
				break
			}
			r++
		}

		numStr := str[digitStart:r]
		n, err := strconv.Atoi(numStr)
		if err != nil {
			return ""
		}

		_, isDigit = ints[rune(str[l])]
		if !isDigit {
			for i := 1; i < n; i++ {
				res = append(res, str[l])
			}
		}
		l = r
	}
	return string(res)
}
