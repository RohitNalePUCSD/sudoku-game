package main

import (
	"fmt"
	"math/rand"
)

//type casting 2d array in matrix
type matrix [9][9]int

var (
	total = 0
	level = 0
	userLevel = 30   			// its hard = 30  medium = 50 easy = 70
	ansGrid = matrix {}			// Anser of grid 
	quegrid = matrix {}			// user Display of grid
)

//genrate random number 
func init() {
	rand.Seed(1500909006430687579)
}

//display the output of grid which is fill all the values
func outputGrid()  {

	for i:= 0; i < 9; i+=1{
		for j:= 0; j< 9; j+=1{
			fmt.Printf("%d\t",ansGrid[i][j])
		}
		fmt.Println()
	}
}

//display the user grid add zeros
func questionGrid(grid matrix)  {
	
	for i:= 0; i < 9; i+=1{
		for j:= 0; j< 9; j+=1{
			fmt.Printf("%d\t",grid[i][j])
			quegrid[i][j] = grid[i][j]
		}
		fmt.Println()
	}

}

// complete grid of value copy into ansGrid
func copyAnswerGrid(a matrix)   {
	
	for i := 0; i < 9; i+=1{
        for j := 0; j < 9; j+=1{
            ansGrid[i][j] = a[i][j];
        }
    }
}

// check its complete grid or not
func completeGrid(a matrix) bool {

    for i := 0; i < 9; i+=1{
        for j := 0; j < 9; j+=1 {
            if(a[i][j] == 0) {
                return false
            }
        }
    }
    return true
}

// check row, col and box (3*3) which is unique value 
func fitGrid(a matrix, value int, row int, col int) bool{

	for i:= 0; i < 9; i++ {
		//check row wise element present or not
		if(a[row][i] == value){
			return false
		}		//check column wise element present or not
		if(a[i][col] == value) {
			return false
		}
	}

	minRow := (row/3)*3;
	maxRow := minRow + 3;
	minCol := (col/3)*3;
	maxCol := minCol + 3;

	//check box (3*3), element present or not
	for i := minRow; i < maxRow; i+=1{
		for j := minCol; j < maxCol; j+=1{
			if(a[i][j] == value){
				return false;
			}
		}
	}
	return true
} 

// fill the grid by using recursive funtions
func fillGrid(grid matrix, row int, col int) int {
	
	//terminat condition of recursionfunction
	if (total == 1){
		return 1
	}


	a := matrix {}
	//copy value previous grid into new genrated grid
	for i:= 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			a[i][j] = grid[i][j]
		}
	}

	//teminated condition, row is full
	if (row == 9) {
		if (completeGrid(a)) {					// check grid is complete or not
			total = 1
			copyAnswerGrid(a)
			return 1
		}
	}
	//fmt.Println("row----------------->",row)
	//fmt.Println("col----------------->",col)

	if(col == 9) {								//check column is completed then move to next row 
		return fillGrid(a, row+1, 0)
	}


	if(a[row][col] == 0) {

		//genrate the 1 number and check is correct or not 
		for i := 1; i <= 9; i++ {	
			if(fitGrid(a, i, row, col))	{
				//put value if its correct 
				a[row][col] = i
				if(total > 0){
					return 1
				}
				fillGrid(a, row, col+1)
			}
		}
	}else{
		return fillGrid(a, row, col+1)
	}
	return 1
}

//add the user level of complexity
func gridGenration(grid matrix)  {
	
	fillGrid(grid, 0, 0)

	//check user  level of the game 
	for ((total > 0) && (level < userLevel)){
		total = 0
		//genrate random number of row and col  
		row := rand.Intn(9)%9
		col := rand.Intn(9)%9
		number := (rand.Intn(9)%9) + 1
		
		//fmt.Println(total, row, col, number)
		//fmt.Println(grid)
		if(grid[row][col] == 0){

			for (fitGrid(grid, number, row, col) == false){
				//genrate random number to satified all rules of sudoku
				number = (rand.Intn(9) %9) +1
			}

			//fmt.Println("fet in gred ", number)
			grid[row][col] = number
			fillGrid(grid, 0, 0)
			//fmt.Println("Run")
			//fmt.Println(grid)

			if(total == 0){
				grid[row][col] = 0;
				total = 1
				continue
			}
			level+=1
		}
		total = 1
	}
	fmt.Println("Question of grid")
	questionGrid(grid)
	fmt.Println("Answer of grid")
	outputGrid()
}

// main fnction
func main()  {

	var grid = matrix {}
	for i:= 0; i< 9; i++{
		for j := 0; j< 9; j++ {
			grid[i][j] = 0
		}
	}
	gridGenration(grid)			//grid genrated function
	userDisplay(quegrid)		// display user grid function
	//userInputGrid(grid)
	//validationGame(grid)
}