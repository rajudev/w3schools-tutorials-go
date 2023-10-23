package main
import ("fmt")

func main() {
	arr1 := [6]int{9,10,11,12,13,14} // an array
	myslice1 := arr1[1:5] // sliced array

	fmt.Println("myslice1 = %v\n", myslice1)
	fmt.Printf("length of myslice1 = %d\n", len(myslice1))
	fmt.Printf("Capacity of myslice1 = %d\n", cap(myslice1))

	myslice1 = arr1[1:3] // change length by reslicing the array
	fmt.Printf("myslice1 = %v \n", myslice1)
	fmt.Printf("length of myslice1 = %d\n", len(myslice1))
	fmt.Printf("Capacity of myslice1 = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v \n", myslice1)
	fmt.Printf("length of myslice1 = %d \n", len(myslice1))
	fmt.Printf("Capacity of myslice1 = %d \n", cap(myslice1))
	
}
