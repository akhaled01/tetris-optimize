package tetops

import (
	"bufio"
	"log"
	"os"
)

func ParseTetrominoesFromFile(filePath string) ([]Tetromino, error) {
	file, err := os.Open(filePath)
	var emptylinecount int
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var tetrominoes []Tetromino
	var tetrominoShape [][]bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			emptylinecount++
			if len(tetrominoShape) > 0 {
				tetrominoes = append(tetrominoes, Tetromino{Shape: tetrominoShape})
				tetrominoShape = nil
			}
		} else {
			var row []bool
			for _, ch := range line {
				if ch == '#' {
					row = append(row, true)
				} else if ch == '.' {
					row = append(row, false)
				} else {
					log.Fatal("ERROR")
				}
			}
			tetrominoShape = append(tetrominoShape, row)
		}
	}
	if len(tetrominoShape) > 0 {
		tetrominoes = append(tetrominoes, Tetromino{Shape: tetrominoShape})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tetrominoes, nil
}
