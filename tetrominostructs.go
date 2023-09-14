package tetops

type Tetromino struct {
	Shape [][]bool
}

type Solver struct {
	Tetrominoes []Tetromino
	Solution    [][]bool
	Size        int
}
