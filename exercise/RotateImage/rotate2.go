package main

/*
You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).
*/

import (
	"fmt"
)

func main() {
	test()
}

func rotate(matrix [][]int) {
	s := 0
    n := len(matrix)
	e := n - 1
	for s < e {
		temp := matrix[s]
		matrix[s] = matrix[e]
		matrix[e] = temp
		s++
		e--
	}

	for i := 0; i < n ; i++ {
		for j := i + 1; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}

func test() {
	m := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	dump(m)

	rotate(m)

	dump(m)

}

func dump(m [][]int) {
	fmt.Println()

	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}
