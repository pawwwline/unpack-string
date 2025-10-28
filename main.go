package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	str, err := unpackString("132")
	fmt.Println(str)
	fmt.Println(err)
}

func unpackString(str string) (string, error) {
	res := make([]rune, 0)
	escape := false
	strRune := []rune(str)
	for i := 0; i < len(strRune); i++ {
		if strRune[i] == '\\' {
			escape = true
			continue
		}

		if !isDigit(strRune[i]) || escape {
			res = append(res, strRune[i])
			escape = false
			numStart := i + 1
			j := numStart

			for j < len(strRune) && isDigit(strRune[j]) {
				j++
			}

			if numStart < j {
				n, err := strconv.Atoi(string(strRune[numStart:j]))
				if err != nil {
					return "", err
				}
				for k := 1; k < n; k++ {
					res = append(res, strRune[numStart-1])
				}
			}
			i = j - 1
		} else {
			return "", errors.New("invalid string")
		}
	}
	return string(res), nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
