package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const INPUT_FILE = "../input.txt"

func getSymbols(lines []string) string {
	const notSymbols = "0123456789."
	validSymbols := ""

	for _, line := range lines {
		for _, lineChar := range line {
			charIsSymbol := true

			// Check whether the current character is a symbol
			for _, nonSymbol := range notSymbols {
				if lineChar == nonSymbol {
					charIsSymbol = false
					break
				}
			}

			// If the current character is a symbol then check if it's already in the list of symbols,
			// if it isn't then add it to the list
			if charIsSymbol {
				uniqueSymbol := true
				
				for _, validSymbol := range validSymbols {
					if validSymbol == lineChar {
						uniqueSymbol = false
						break
					}
				}

				if uniqueSymbol {
					validSymbols = validSymbols + string(lineChar)
				}
			}
		}
	}

	return validSymbols
}

func isCharValidInt (char rune) bool {
	const validInts = "0123456789"

	for _, validInt := range validInts {
		if char == validInt {
			return true
		}
	}
	return false
}

func isCharValidSymbol (char rune, symbols string) bool {
	isValid := false

	for _, symbol := range symbols {
		if char == symbol {
			isValid = true
		}
	}

	return isValid
}

func main() {
	file, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	validSymbols := getSymbols(lines)

	enginePartNumbers := []int{}

	// Iterate over all lines
	for lineIndex, line := range lines {
		// For each line iterate over all characters
		for lineCharIndex := 0; lineCharIndex < len(line); lineCharIndex++ {
			lineChar := line[lineCharIndex]

			if isCharValidInt(rune(lineChar)) {
				enginePartNrIndexStart := lineCharIndex
				enginePartNrIndexEnd := -1

				// If current character in the line is an integer then iterate to find the end of the integer
				for i := enginePartNrIndexStart; i < len(line); i++ {
					// If the last character of the line is a valid int then set it as the engine nr index end
					if i + 1 == len(line) && isCharValidInt(rune(line[i])) {
						enginePartNrIndexEnd = i
						break
					} else if !isCharValidInt(rune(line[i])) {
						enginePartNrIndexEnd = i - 1
						break
					}
				}

				if enginePartNrIndexEnd == -1 {
					log.Fatal("Failed finding the end index of an engine number")
				}
				
				// Skip to index right after where the current engine number ends
				lineCharIndex = enginePartNrIndexEnd + 1

				engineNr := line[enginePartNrIndexStart:enginePartNrIndexEnd+1]

				// Check that the engine number is a valid one by checking whether it's surrounded by symbols
				isValidEngineNumber := false

				// Check on the line above
				isLineAbove := lineIndex > 0
				if isLineAbove {
					aboveLine := lines[lineIndex - 1]
					leftBound := enginePartNrIndexStart - 1
					rightBound := enginePartNrIndexEnd + 2

					if enginePartNrIndexStart == 0 {
						leftBound = 0
					}

					if enginePartNrIndexEnd + 1 == len(aboveLine) {
						rightBound = enginePartNrIndexEnd + 1
					}

					for i := leftBound; i < rightBound; i++ {
						if isCharValidSymbol(rune(aboveLine[i]), validSymbols) {
							isValidEngineNumber = true
							break
						}
					}
				}

				// Check left
				isSpaceLeft := enginePartNrIndexStart > 0
				if isSpaceLeft && !isValidEngineNumber {
					if isCharValidSymbol(rune(line[enginePartNrIndexStart-1]), validSymbols) {
						isValidEngineNumber = true
					}
				}

				// Check right
				isSpaceRight := enginePartNrIndexEnd + 1 < len(line)
				if isSpaceRight && !isValidEngineNumber {
					if isCharValidSymbol(rune(line[enginePartNrIndexEnd + 1]), validSymbols) {
						isValidEngineNumber = true
					}
				}

				// Check on the line below
				isLineBelow := lineIndex + 1 < len(lines)
				if isLineBelow && !isValidEngineNumber {
					belowLine := lines[lineIndex + 1]
					leftBound := enginePartNrIndexStart - 1
					rightBound := enginePartNrIndexEnd + 2

					if enginePartNrIndexStart == 0 {
						leftBound = 0
					}

					if enginePartNrIndexEnd + 1 == len(belowLine) {
						rightBound = enginePartNrIndexEnd + 1
					}

					for i := leftBound; i < rightBound; i++ {
						if isCharValidSymbol(rune(belowLine[i]), validSymbols) {
							isValidEngineNumber = true
							break
						}
					}
				}

				if isValidEngineNumber {
					engineNrInt, err := strconv.Atoi(engineNr)
					if err != nil {
						log.Fatal("Failed converting engine number string to integer:", err)
					}

					enginePartNumbers = append(enginePartNumbers, engineNrInt)
				}
			}
		}
	}

	enginePartNrSum := 0
	for _, enginePartNr := range enginePartNumbers {
		enginePartNrSum += enginePartNr
	}
	fmt.Println(enginePartNrSum)
}