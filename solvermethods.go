package tetops

import "math"

func (s *Solver) ReArrangeTetrominoes() {
	for index := 0; index < len(s.Tetrominoes); index++ {
		startRow := FindStartingRow(s.Tetrominoes[index].Shape)
		startCol := FindStartingCol(s.Tetrominoes[index].Shape)
		if (startRow == 0 && startCol == 0) || startRow == -1 || startCol == -1 {
			continue
		}
		// Shift the tetromino by updating the values in the slice
		for i := startRow; i < len(s.Tetrominoes[index].Shape); i++ {
			for j := startCol; j < len(s.Tetrominoes[index].Shape); j++ {
				s.Tetrominoes[index].Shape[i-startRow][j-startCol] = s.Tetrominoes[index].Shape[i][j]
				s.Tetrominoes[index].Shape[i][j] = false
			}
		}
	}
}

// Find the starting row of the tetromino
func FindStartingRow(slice [][]bool) int {
	for rowIndex := range slice {
		for colIndex := range slice[rowIndex] {
			if slice[rowIndex][colIndex] {
				return rowIndex
			}
		}
	}
	return -1 // If no starting row is found, return -1 as an invalid value
}

// Find the starting column of the tetromino
func FindStartingCol(slice [][]bool) int {
	for colIndex := range slice {
		for rowIndex := range slice[colIndex] {
			if slice[rowIndex][colIndex] {
				return colIndex
			}
		}
	}
	return -1 // If no starting column is found, return -1 as an invalid value
}

func (s *Solver) placeTetromino(index int) bool {
	if index == len(s.Tetrominoes) {
		return true
	}

	for row := 0; row < s.Size; row++ {
		for col := 0; col < s.Size; col++ {
			if s.canPlaceTetromino(index, row, col) {
				s.addTetromino(index, row, col)

				if s.placeTetromino(index + 1) {
					return true
				}

				s.removeTetromino(index, row, col)
			}
		}
	}

	return false
}

func (s *Solver) canPlaceTetromino(index, row, col int) bool {
	tetromino := s.Tetrominoes[index]

	for i, tRow := range tetromino.Shape {
		for j, cell := range tRow {
			if cell && (row+i >= s.Size || col+j >= s.Size || s.Solution[row+i][col+j] != '.') {
				return false
			}
		}
	}

	return true
}


func (s *Solver) addTetromino(index, row, col int) {
	tetromino := s.Tetrominoes[index]

	for i, tRow := range tetromino.Shape {
		for j, cell := range tRow {
			if cell {
				s.Solution[row+i][col+j] = rune(index+'A')
			}
		}
	}
}

func (s *Solver) removeTetromino(index, row, col int) {
	tetromino := s.Tetrominoes[index]

	for i, tRow := range tetromino.Shape {
		for j, cell := range tRow {
			if cell {
				s.Solution[row+i][col+j] = '.'
			}
		}
	}
}

func NewSolver(tetrominoes []Tetromino) *Solver {
	return &Solver{
		Tetrominoes: tetrominoes,
		Solution:    nil,
		Size:        0,
	}
}

func (s *Solver) Solve() bool {
	s.Size = int(math.Sqrt(float64(len(s.Tetrominoes) * 4)))
	for {
		s.Solution = make([][]rune, s.Size)
		for i := range s.Solution {
			s.Solution[i] = make([]rune, s.Size)
		}

		for v := range s.Solution {
			for j := range s.Solution[v] {
				s.Solution[v][j] = '.'
			}
		}

		if s.placeTetromino(0) {
			return true
		}

		s.Size++
	}
}
