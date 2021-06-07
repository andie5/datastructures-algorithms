package main

import "log"

// https://www.hackerrank.com/challenges/magic-square-forming/problem

// Magic squares

// Sum of any row, column, or diagonal of length is
// always equal to same number
// Distinct integers 1-9
// Input 3x3 array
// Print cost (each swap made) & add up
// Goal: smallest cost for given input -> magic square

// E.g.
// $s = [[5, 3, 4], [1, 5, 8], [6, 4, 2]]
// 5 3 4
// 1 5 8
// 6 4 2

// Add up rows, columns and diagonals to find common sum
// row 0 = 12 -> 3      [0][0] 5-> 8
// row 1 = 14 -> 1		[1][2] 8 -> 9
// row 2 = 12 -> 3

// col 0 = 12 -> 3		[0][0] 5-> 8
// col 1 = 12 -> 3
// col 2 = 14 -> 1		[1][2] -> 9

// diag 1 \ = 12 -> 3	[0][0] 5-> 8
// diag 2 / = 15

// 8 3 4
// 1 5 9
// 6 4 2

// Common sum = 12

// 4 and 5 = 9
// Need 7, 9  = 15

// ----------

// 4 9 2
// 3 5 7
// 8 1 5

// row 0 = 15
// row 1 = 15
// row 2 = 14

// col 0 = 15
// col 1 = 15
// col 2 = 14

// diag 1 \ = 14
// diag 2 / = 15
// common sum 15

// Need 6
// --------

// 4 8 2
// 4 5 7
// 6 1 6

// row 0 = 14
// row 1 = 16
// row 2 = 13
// col 0 = 14
// col 1 = 14
// col 2 = 15
// diag 1 \ = 15
// diag 2 / = 13
// Common sum = 14

// need 9, 3

// Steps
// Find magic constant - formula
// find what number needs to be added to each row  col, diag to make up total
// Find which distinct number is missing
// and any duplicatesin square
// See which squares can change

// Find magic constanst 3x3 square
// 1+2+3+4+5+6+7+8+9+10+11+12+13+14+15=45
// magic = 45/3 = 15

func findMagicConstant(s [][]int32) int {
	sum := 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9
	matrix := len(s)
	return sum / matrix
}

func getSumRowColumnDiag(s [][]int32) (map[int]int32, map[int]int32, map[int]int32) {
	rows := map[int]int32{}
	cols := map[int]int32{}
	diag := map[int]int32{}

	for i, row := range s {
		for j, item := range row {
			rows[i] += row[j]
			cols[j] += item
		}
	}
	diag[0] = s[0][0] + s[1][1] + s[2][2]
	diag[1] = s[0][2] + s[1][1] + s[2][0]
	return rows, cols, diag
}

// func getDistinctMissingNums(s [][]int32) (map[int]int32, []int32) {
// 	x := map[int]int32{}
// 	for i := 1; i < 10; i++ {
// 		x[i] = int32(0)
// 	}

// 		for rowI, row := range s {
// 			for j, item := range row{
// 				 x[int(item)] += 1
// 				 if x[int(item)]
// 			}
// 		}
// 	return x
// }

func formingMagicSquare(s [][]int32) int32 {
	magicConstant := findMagicConstant(s)
	rows, columns, diag := getSumRowColumnDiag(s)
	// missing := getDistinctMissingNums(s)

	log.Println("rows: ", rows)
	log.Println("columns: ", columns)
	log.Println("diag: ", diag)
	log.Println("magicConstant: ", magicConstant)
	return int32(0)
}

func main() {
	sample1 := [][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}
	formingMagicSquare(sample1)
}
