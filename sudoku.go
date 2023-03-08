package main

import (
	"os"

	"github.com/01-edu/z01"
)

func print(arr [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(rune(arr[i][j] + 48))
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

func isSafe(grid [][]int, row, col, num int) bool {
	for x := 0; x < 9; x++ {
		if grid[row][x] == num {
			return false
		}
	}
	for x := 0; x < 9; x++ {
		if grid[x][col] == num {
			return false
		}
	}
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

func solveSudoku(grid [][]int, row, col int) bool {
	if row == 8 && col == 9 {
		return true
	}
	if col == 9 {
		row++
		col = 0
	}
	if grid[row][col] > 0 {
		return solveSudoku(grid, row, col+1)
	}
	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num
			if solveSudoku(grid, row, col+1) {
				return true
			}
		}
		grid[row][col] = 0
	}
	return false
}

func main() {
	input := os.Args[1:]
	var grid [][]int
	if len(os.Args[1:]) == 0 {
		z01.PrintRune('E')
		z01.PrintRune('r')
		z01.PrintRune('r')
		z01.PrintRune('o')
		z01.PrintRune('r')
		z01.PrintRune('\n')
	} else {

		grid = toInt(input)
		if solveSudoku(grid, 0, 0) {
			print(grid)
		} else {
			z01.PrintRune('E')
			z01.PrintRune('r')
			z01.PrintRune('r')
			z01.PrintRune('o')
			z01.PrintRune('r')
			z01.PrintRune('\n')
		}
	}
}

func toInt(input []string) [][]int {
	grid := make([][]int, 9)
	for i := range grid {
		grid[i] = make([]int, 9)
	}
	for i, row := range input {
		for j, val := range row {
			if val == '.' {
				grid[i][j] = 0
			} else {
				grid[i][j] = int(val - 48)
			}
		}
	}
	return grid
}
