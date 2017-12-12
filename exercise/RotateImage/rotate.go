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
	n := len(matrix)
	startX := 0
	startY := 0
	size := n
	for round := 0; round < n/2; round++ {

		for i := 0; i < size-1; i++ {
			temp := matrix[startX][startY+i]
			matrix[startX][startY+i] = matrix[startX+(size-1-i)][startY]
			matrix[startX+(size-1-i)][startY] = matrix[startX+(size-1)][startY+(size-1-i)]
			matrix[startX+(size-1)][startY+(size-1-i)] = matrix[startX+i][startY+(size-1)]
			matrix[startX+i][startY+(size-1)] = temp
		}

		startX++
		startY++
		size = size - 2
	}
}

/*
func rotateRound(startX int, startY int, size int, matrix [][]int) {
	for i := 0; i < size-1; i++ {
		temp := matrix[startX][startY+i]
		matrix[startX][startY+i] = matrix[startX+(size-1-i)][startY]
		matrix[startX+(size-1-i)][startY] = matrix[startX+(size-1)][startY+(size-1-i)]
		matrix[startX+(size-1)][startY+(size-1-i)] = matrix[startX+i][startY+(size-1)]
		matrix[startX+i][startY+(size-1)] = temp
	}
}
*/

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
