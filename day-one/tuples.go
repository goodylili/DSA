package main

import "fmt"

func powerSeries(a int) (square, cube int) {
	square = a * a
	cube = square * a
	return square, cube
}

func main() {
	var square, cube int
	square, cube = powerSeries(10)
	fmt.Println("Square is ", square, "Cube is ", cube)
}
