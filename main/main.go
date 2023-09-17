package main

import (
	"fmt"
	"log"
	"os"
	"tetops"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("[USAGE]: go run . <tetromino_file>")
		os.Exit(1)
	}

	tetrominoes, err := tetops.ParseTetrominoesFromFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	solver := tetops.NewSolver(tetrominoes)
	solver.ReArrangeTetrominoes()
	if solver.Solve() {
		fmt.Printf("Smallest square size: %d\n", solver.Size)
		fmt.Println("Solution:")
		for _, row := range solver.Solution {
			for _, cell := range row {
				if cell != '.' {
					fmt.Print(string(cell)) // Filled cell
				} else {
					fmt.Print(".") // Empty cell
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("No solution found.")
	}
}
