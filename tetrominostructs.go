package tetops

type Tetromino struct {
	Shape [][]bool
}

type Solver struct {
	Tetrominoes []Tetromino
	Solution    [][]rune
	Size        int
}
