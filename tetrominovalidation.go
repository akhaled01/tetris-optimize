package tetops

import (
	"bufio"
	"log"
	"os"
)

func (s *Solver) ValidateTetrominos() bool {
	// loop through each tetromino and check number of # and connections
	hashcount := 0
	connectionCount := 0
	for _, tetromino := range s.Tetrominoes {
		if len(tetromino.Shape) != 4 {
			return false
		}
		for i, row := range tetromino.Shape {
			if len(row) != 4 {
				return false
			}
			for j, ch := range row {
				if ch {
					hashcount++
					if i < 3 && tetromino.Shape[i+1][j] {
						connectionCount++
					}
					if i > 0 && tetromino.Shape[i-1][j] {
						connectionCount++
					}
					if j < 3 && tetromino.Shape[i][j+1] {
						connectionCount++
					}
					if j > 0 && tetromino.Shape[i][j-1] {
						connectionCount++
					}
				}
			}
		}
	}
	return connectionCount >= 6*len(s.Tetrominoes) && hashcount == 4*len(s.Tetrominoes)
}

func CheckEmptyLineFollowedByEmptyLine(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		return 
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isEmptyLine := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if isEmptyLine {
				log.Fatal("ERROR")
			}
			isEmptyLine = true
		} else {
			isEmptyLine = false
		}
	}

	if err := scanner.Err(); err != nil {
		return 
	}
}
