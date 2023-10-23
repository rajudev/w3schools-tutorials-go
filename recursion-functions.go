// Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

package main
import ("fmt")

func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x+1)
}

func main() {
	testCount(1)
}
