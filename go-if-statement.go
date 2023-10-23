// Use the if statement to specify a block of Go code to be executed if a condition is true.

package main
import ("fmt")

func main() {
	x := 10
	y := 3

	if x > y {
		fmt.Printf("%v is greater than %v \n", x, y)
	}
}
