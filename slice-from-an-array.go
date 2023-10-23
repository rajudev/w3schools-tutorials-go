package main
import ("fmt")

func main() {
	arr1 := [6]int{1,2,3,4,5,6}
	myslice := arr1[1:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
	

}
