package main
import ("fmt")

func main() {
	myslice1 := []int{1,2,3}
	myslice2 := []int{4,5,6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Println("myslice3 = %v\n", myslice3)
	fmt.Println("Length of myslice3 = %d\n", len(myslice3))
	fmt.Println("capacity of myslice3 = %n", cap(myslice3))
		
}
