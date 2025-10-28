package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(str), nil
}

func main() {
	str, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	unpacked, err := unpackString(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unpacked)
}

func unpackString(str string) (string, error) {
	var b strings.Builder
	escape := false
	strRune := []rune(str)
	for i := 0; i < len(strRune); i++ {
		if strRune[i] == '\\' {
			escape = true
			continue
		}

		if !isDigit(strRune[i]) || escape {
			escape = false
			numStart := i + 1
			j := numStart

			if numStart < len(strRune) && strRune[numStart] == '0' {
				i = numStart
				continue
			}
			b.WriteRune(strRune[i])

			for j < len(strRune) && isDigit(strRune[j]) {
				j++
			}

			if numStart < j {
				n, err := strconv.Atoi(string(strRune[numStart:j]))
				if err != nil {
					return "", err
				}
				for k := 1; k < n; k++ {
					b.WriteRune(strRune[numStart-1])
				}
			}
			i = j - 1
		} else {
			return "", errors.New("invalid string")
		}
	}
	return b.String(), nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
