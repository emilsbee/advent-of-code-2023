package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const INPUT_FILE = "../input.txt"

func isCharValidInt (char rune) bool {
	const validInts = "0123456789"

	for _, validInt := range validInts {
		if char == validInt {
			return true
		}
	}
	return false
}

func isCharAGear (char rune) bool {
	return char == '*'
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

	gearRatios := []int{}

	// Iterate over all lines
	for lineIndex, line := range lines {
		// For each line iterate over all characters
		for lineCharIndex := 0; lineCharIndex < len(line); lineCharIndex++ {
			lineChar := line[lineCharIndex]

			if isCharAGear(rune(lineChar)) {
				gearIndex := lineCharIndex
				gearNumbers := []int{}

				// Check on the right
				isSpaceRight := gearIndex + 1 < len(line)
				if isSpaceRight {
					if isCharValidInt(rune(line[gearIndex + 1])) {
						numberEnd := -1

						// Find where the number ends
						for i := gearIndex + 1; i + 1 < len(line); i++ {
							if !isCharValidInt(rune(line[i])) {
								numberEnd = i - 1
								break
							} else if i + 1 == len(line) {
								numberEnd = i
							}
						}

						if numberEnd == -1 {
							log.Fatal("Failed finding number end on the right")
						}

						gearNumber, _ := strconv.Atoi(line[gearIndex+1:numberEnd+1])

						// Save the found number
						gearNumbers = append(gearNumbers, gearNumber)
					}
				}

				// Check on the left
				isSpaceLeft := gearIndex > 0
				if isSpaceLeft {
					if isCharValidInt(rune(line[gearIndex - 1])) {
						numberBegin := -1

						// Find where the number begins
						for i := gearIndex - 1; i >= 0; i-- {
							if !isCharValidInt(rune(line[i])) {
								numberBegin = i + 1
								break
							} else if i == 0 {
								numberBegin = 0
							}
						}

						if numberBegin == -1 {
							log.Fatal("Failed finding number beginning on the left")
						}

						gearNumber, _ := strconv.Atoi(line[numberBegin:gearIndex])

						// Save the found number
						gearNumbers = append(gearNumbers, gearNumber)
					}
				}

				// Check below
				isLineBelow := lineIndex + 1 < len(lines)
				if isLineBelow {
					belowLine := lines[lineIndex + 1]
					leftBound := gearIndex - 1
					rightBound := gearIndex + 2

					if gearIndex == 0 {
						leftBound = 0
					}

					if gearIndex + 1 == len(belowLine) {
						rightBound = gearIndex + 1
					}

					for i := leftBound; i < rightBound; i++ {
						if isCharValidInt(rune(belowLine[i])) {
							numberBegin := -1
							numberEnd := -1

							// Find index where the number begins
							for j := i; j >= 0; j-- {
								if !isCharValidInt(rune(belowLine[j])) {
									numberBegin = j + 1
									break
								} else if j == 0 {
									numberBegin = 0
								}
							}

							if numberBegin == -1 {
								log.Fatal("Could not find where number begins for below line")
							}

							// Find where the number ends
							for k := numberBegin; k < len(belowLine); k++ {
								if !isCharValidInt((rune(belowLine[k]))) {
									numberEnd = k - 1
									break
								} else if k + 1 == len(belowLine) {
									numberEnd = k
								}
							}

							if numberEnd == -1 {
								log.Fatal("Could not find where number ends for below line")
							}

							// Reset the loop that iterates "line below" to index after the found number
							if numberEnd + 1 < rightBound {
								i = numberEnd + 1
							} else if numberEnd + 1 == rightBound {
								i = rightBound
							} else if numberEnd + 1 == len(belowLine) {
								i = rightBound
							} else if numberEnd == rightBound {
								i = rightBound
							}

							foundNumber, _ := strconv.Atoi(belowLine[numberBegin:numberEnd+1])	

							// Save the found number
							gearNumbers = append(gearNumbers, foundNumber)
						}
					}
				}

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
							numberBegin := -1
							numberEnd := -1

							// Find index where the number begins
							for j := i; j >= 0; j-- {
								if !isCharValidInt(rune(aboveLine[j])) {
									numberBegin = j + 1
									break
								} else if j == 0 {
									numberBegin = j
								}
							}

							if numberBegin == -1 {
								log.Fatal("Could not find where number begins for above line")
							}

							// Find index where the number ends
							for k := numberBegin; k < len(aboveLine); k++ {
								if !isCharValidInt((rune(aboveLine[k]))) {
									numberEnd = k - 1
									break
								} else if k + 1 == len(aboveLine) {
									numberEnd = k
								}
							}

							if numberEnd == -1 {
								log.Fatal("Could not find where number ends for above line")
							}

							// Reset the loop that iterates "line above" to index after the found number
							if numberEnd + 1 < rightBound {
								i = numberEnd + 1
							} else if numberEnd + 1 == rightBound {
								i = rightBound
							} else if numberEnd + 1 == len(aboveLine) {
								i = rightBound
							} else if numberEnd == rightBound {
								i = rightBound
							}

							foundNumber, _ := strconv.Atoi(aboveLine[numberBegin:numberEnd+1])	

							// Save the found number
							gearNumbers = append(gearNumbers, foundNumber)
						}
					}
				}

				if len(gearNumbers) == 2 {
					gearRatio := gearNumbers[0] * gearNumbers[1]
					gearRatios = append(gearRatios, gearRatio)
				}
			}
		}
	}

	gearRatioSum := 0
	for _, gearRatio := range gearRatios {
		gearRatioSum += gearRatio
	}
	fmt.Println(gearRatioSum)
}