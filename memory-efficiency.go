/*
   When using slices, Go loads all the underlying elements into the memory.

   If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.

   The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program. 
   Syntax copy(dest, src)

   The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.
   Example

   This example shows how to use the copy() function:
*/

package main
import ("fmt")

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	//Original Slice
	fmt.Printf("numbers = %v \n", numbers)
	fmt.Printf("length of numbers = %d \n", len(numbers))
	fmt.Printf("Capacity of numbers = %d \n", cap(numbers))

	// Create a copy of only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numberscopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("Capacity = %d\n", cap(numbersCopy))
		
}
