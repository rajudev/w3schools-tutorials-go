package main
import ("fmt")

func main() {

	myslice1 := []int{1,2,3,4,5,6}

	fmt.Println("myslice1 = %v\n", myslice1)
	fmt.Println("length of myslice1 = %d\n", len(myslice1))
	fmt.Println("Capacity of myslice1 = %d\n", cap(myslice1))

	// Appending more elements to myslice1
	myslice1 = append(myslice1, 20, 21)
	
	fmt.Println("myslice1 = %v\n", myslice1)
	fmt.Println("length of myslice1 = %d\n", len(myslice1))
	fmt.Println("Capacity of myslice1 = %d\n", cap(myslice1))

	

}
