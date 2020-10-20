package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Array of points on a grid for the robot to travel in
// Uses 2 times the x and y axis for offsetting of negative points
var array [2 * k_x_axis_max][2 * k_y_axis_max]Point

// Variable global count for when the recursive CheckPoints
// function loops it will count up all safe points for the robot to move on
var count int

func main() {
	// Count starts at 1 due to the first position the 
	// robot tries is (0, 0) which is safe
	// If count is 0 then it will miss its first point
	count = 1
	array[k_x_axis_max][k_y_axis_max].Visited = true
	CheckPoints(0, 0)
	fmt.Println("Total Area Points: " + strconv.Itoa(count))
}

func CheckPoints(x int, y int) {

	// Hard coded movements that the robot can move in
	// In the future Id like to contain this a struct of
	// struct to edit the movements that can be performed
	array_movements := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, b := range array_movements {

		next_move_x := b[0]
		next_move_y := b[1]

		// Get sum of the abs next move
		sum := GetSum(x+next_move_x,
			y+next_move_y)

		// Is it safe for the sum calculated
		safe := IsSafe(sum)

		// Were the next move will be, offset by the max
		// to allow for negative numbers
		x_visit := x + b[0] + k_x_axis_max
		y_visit := y + b[1] + k_y_axis_max

		// Ensures if a incorrect index is found in this recursive loop 
		// It must exit the program
		if x_visit < 0 {
			fmt.Println("Error indexing x axis")
			os.Exit(1)
		} else if y_visit < 0 {
			fmt.Println("Error indexing y axis")
			os.Exit(1)
		}

		// Checks whether safe is true and if we have already visited this point
		if safe && !array[x_visit][y_visit].Visited {
			count++
			// Puts the offset value as visited again to ensure that if a negative index is used
			// Then a absolute sum would still work but
			// as we can't index in a golang array with negatives
			// this a work around
			array[x+k_x_axis_max+b[0]][y+k_y_axis_max+b[1]].Visited = true
			// Calls itself again recursively with the next move that is safe and
			// has not been visited yet
			CheckPoints(x+b[0], y+b[1])
		}
	}
}


// Converts a int to float64 for absolute checking
// Splits a number e.g. 69 into 6 and 9
// By using the modulus and then creating a running sum
// of all its individual digits
func FilterDigits(number int) int {
	abs_number := int(math.Abs(float64(number)))
	sum := 0
	for abs_number > 0 {
		sum += abs_number % 10
		abs_number = abs_number / 10
	}
	return sum
}


// Places x and y into the FilterDigits function
// to create a total value sum
func GetSum(x1 int, y1 int) int {
	total := FilterDigits(x1) + FilterDigits(y1)
	return total
}


// Checker fucntion to know when a total value
// is less than the max safe value
func IsSafe(total int) bool {
	if total <= k_safe_value {
		return true
	} else {
		return false
	}
}
