package main

import "fmt"

func main() {

	var places [10][10]string
	var row, col int
	fmt.Printf("Enter no of rows:\n")
	fmt.Scanf("%d\n", &row)
	fmt.Printf("Enter no of columns :\n")
	fmt.Scanf("%d\n", &col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" food %d %d :\t", i+1, j+1)
			fmt.Scanf("%s\t", &places[i][j])
		}
	}
	fmt.Printf("\t")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%s\t", places[i][j])
			if j == col-1 {
				fmt.Printf(" \t ")

			}
		}
	}

}
