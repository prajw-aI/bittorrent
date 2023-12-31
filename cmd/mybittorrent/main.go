package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func decodeBencode(bencodedString string) (interface{}, error) {
	if unicode.IsDigit(rune(bencodedString[0])) {
		var firstColonIndex int

		for i := 0; i < len(bencodedString); i++ {
			if bencodedString[i] == ':' {
				firstColonIndex = i
				break
			}
		}

		lengthStr := bencodedString[:firstColonIndex]
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return "", err
		}
		return bencodedString[firstColonIndex+1 : firstColonIndex+1+length], err
	} else if bencodedString[0] == 'i' && bencodedString[len(bencodedString)-1] == 'e' {
		numstr, err := strconv.Atoi(bencodedString[1 : len(bencodedString)-1])
		if err != nil {
			return "", err
		}
		return numstr, nil
	} else {
		return "", fmt.Errorf("only strings are supported at the moment")
	}
}

func main() {
	command := os.Args[1]
	if command == "decode" {
		bencodedValue := os.Args[2]
		decoded, err := decodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}
		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
