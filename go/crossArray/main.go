// Fill in with 1's the row and column for each matrix location having the 1 value.

/*
  Input array
  [1 0 0]
  [0 0 0]
  [0 0 1]

  Resulting array
  [1 1 1]
  [1 0 1]
  [1 1 1]
*/

package main

import (
	"fmt"
	"math/rand"
)

// newRandomMatrix generates some random content.
func newRandomMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)
	for i, _ := range m {
		m[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			// Increase the chance of having zero filled lines/columns.
			if rand.Intn(10) > 8 {
				m[i][j] = 1
			}
		}
	}
	return m
}

// printMatrix prints the matrix
func printMatrix(m [][]int) {
	for _, l := range m {
		fmt.Println(l)
	}
	fmt.Println()
}

// matrixExercise runs the exercise for a given matrix
func matrixExercise(m [][]int) {
	var rows []int
	var lines []int

	fmt.Println("Input array")
	printMatrix(m)

	t := m
	for i, l := range m {
		for j, e := range l {
			if e == 1 {
				rows = append(rows, i)
				lines = append(lines, j)
			}
		}
	}

	for i, l := range m {
		for j, _ := range l {
			for _, v := range rows {
				if v == i {
					m[i][j] = 1
				}
			}
			for _, v := range lines {
				if v == j {
					m[i][j] = 1
				}
			}
		}
	}

	fmt.Println("Resulting array")
	printMatrix(t)
}

func main() {
	m := [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	matrixExercise(m)

	m = newRandomMatrix(5, 7)
	matrixExercise(m)
}
