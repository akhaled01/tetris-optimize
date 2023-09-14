package tetops

import "math"

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
			if cell && (row+i >= s.Size || col+j >= s.Size || s.Solution[row+i][col+j]) {
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
				s.Solution[row+i][col+j] = true
			}
		}
	}
}

func (s *Solver) removeTetromino(index, row, col int) {
	tetromino := s.Tetrominoes[index]

	for i, tRow := range tetromino.Shape {
		for j, cell := range tRow {
			if cell {
				s.Solution[row+i][col+j] = false
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
		s.Solution = make([][]bool, s.Size)
		for i := range s.Solution {
			s.Solution[i] = make([]bool, s.Size)
		}

		if s.placeTetromino(0) {
			return true
		}

		s.Size++
	}
}
