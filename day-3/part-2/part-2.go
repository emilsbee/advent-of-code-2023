package main

import (
	"bufio"
	"log"
	"os"
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

func isCharValidGear (char rune) bool {
	return char == '*'
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
 	_ = validSymbols

	gearNumbers := []int{}
	_ = gearNumbers

	// Iterate over all lines
	for lineIndex, line := range lines {
		// For each line iterate over all characters
		for lineCharIndex := 0; lineCharIndex < len(line); lineCharIndex++ {
			lineChar := line[lineCharIndex]

			if isCharValidGear(rune(lineChar)) {
				gearIndex := lineCharIndex
				isValidGear := false
				gearNumberOne := -1
				gearNumberTwo := -1
				_  = gearNumberOne
				_  = gearNumberTwo


				// Check on the line above
				isLineAbove := lineIndex > 0
				if isLineAbove {
					aboveLine := lines[lineIndex - 1]
					leftBound := gearIndex - 1
					rightBound := gearIndex + 2

					if gearIndex == 0 {
						leftBound = 0
					}

					if gearIndex + 1 == len(aboveLine) {
						rightBound = gearIndex + 1
					}

					for i := leftBound; i < rightBound; i++ {
						if isCharValidInt(rune(aboveLine[i])) {
							endOfInt := i
							// Iterate leftward until we hit a non int character
							for j := leftBound; j < rightBound; j++ {
								// Found the beginning of the int
								if !isCharValidInt(rune(aboveLine[i])) {
									for k := j + 1; k < rightBound; k++ {
										// Found the end of the integer
										if !isCharValidInt((rune(aboveLine[k]))) {
											if k + 1 != rightBound {
												// Reset the iterator that goes through all the chars on above line to the
												// index one after end of the integer we found
												i = k + 1
											}
										}
									}
								}
							}
						}
					}
				}

				// // Check left
				// isSpaceLeft := enginePartNrIndexStart > 0
				// if isSpaceLeft && !isValidEngineNumber {
				// 	if isCharValidSymbol(rune(line[enginePartNrIndexStart-1]), validSymbols) {
				// 		isValidEngineNumber = true
				// 	}
				// }

				// // Check right
				// isSpaceRight := enginePartNrIndexEnd + 1 < len(line)
				// if isSpaceRight && !isValidEngineNumber {
				// 	if isCharValidSymbol(rune(line[enginePartNrIndexEnd + 1]), validSymbols) {
				// 		isValidEngineNumber = true
				// 	}
				// }

				// // Check on the line below
				// isLineBelow := lineIndex + 1 < len(lines)
				// if isLineBelow && !isValidEngineNumber {
				// 	belowLine := lines[lineIndex + 1]
				// 	leftBound := enginePartNrIndexStart - 1
				// 	rightBound := enginePartNrIndexEnd + 2

				// 	if enginePartNrIndexStart == 0 {
				// 		leftBound = 0
				// 	}

				// 	if enginePartNrIndexEnd + 1 == len(belowLine) {
				// 		rightBound = enginePartNrIndexEnd + 1
				// 	}

				// 	for i := leftBound; i < rightBound; i++ {
				// 		if isCharValidSymbol(rune(belowLine[i]), validSymbols) {
				// 			isValidEngineNumber = true
				// 			break
				// 		}
				// 	}
				// }

				// if isValidEngineNumber {
				// 	engineNrInt, err := strconv.Atoi(engineNr)
				// 	if err != nil {
				// 		log.Fatal("Failed converting engine number string to integer:", err)
				// 	}

				// 	gearNumbers = append(gearNumbers, engineNrInt)
				// }
			}
		}
	}

	// gearNrSum := 0
	// for _, gearNr := range gearNumbers {
	// 	gearNrSum += gearNr
	// }
	// fmt.Println(gearNrSum)
}