package main
import ("fmt")

func sum(x int, y int) (result int) {
	result = x + y
	return result
}

func main() {
	fmt.Println(sum(4,34))
	fmt.Println(sum(234,234))
}
