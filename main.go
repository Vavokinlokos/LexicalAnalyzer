package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := "g5ggdfg + jjj 666 -12.03 12 -5 12634576.548461 + *(*# 5"
	input += " "
	fmt.Println("Input is ", input)
	matrix := [8][5]int{
		{1, 4, 3, 0, 0},
		{0, 2, 0, 0, 0},
		{6, 0, 0, 0, 0},
		{0, 4, 0, 0, 0},
		{0, 4, 0, 5, 0},
		{0, 5, 0, 0, 0},
		{6, 0, 0, 0, 0},
	}

	matrixClass := make(map[int]string)
	matrixClass[4] = " -- Integer constant"
	matrixClass[5] = " -- Real constant"
	matrixClass[6] = " -- Identifier"

	resultInput := strings.Split(input, " ")
	resultClasses := make([]string, 0, len(resultInput))
	errorFlag := false
	i, j := 0, 0
	currentIndex := 0
	for currentIndex < len(input) {
		if errorFlag {
			if input[currentIndex] == uint8([]rune(" ")[0]) {
				errorFlag = false
			}
			currentIndex++
			continue
		}
		j = GetJ(int32(input[currentIndex]))
		if j == -1 {
			resultClasses = append(resultClasses, " -- Invalid input")
			errorFlag = true
			currentIndex++
			continue
		}
		prev := i
		i = matrix[i][j]
		if i == 0 {
			if class, ok := matrixClass[prev]; ok {
				resultClasses = append(resultClasses, class)
			} else {
				resultClasses = append(resultClasses, " -- Invalid input")
				errorFlag = true
			}
		}
		currentIndex++

	}
	for i, code := range resultClasses {
		fmt.Println(resultInput[i], code)
	}
	return
}

func GetJ(letter int32) int {
	switch {
	case unicode.IsLetter(letter):
		return 0
	case unicode.IsDigit(letter):
		return 1
	case letter == []rune("-")[0]:
		return 2
	case letter == []rune(".")[0]:
		return 3
	case letter == []rune(" ")[0]:
		return 4
	default:
		return -1
	}
}
