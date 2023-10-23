package main
import ("fmt")

func myFunction(x int, y string) (result int, text1 string) {
	result = x + x
	text1 = y + "World"
	return
}

func main(){
	a, b := myFunction(10, "Hello ")
	fmt.Println(a,b)
}
