//for index, value := array|slice|map

package main
import ("fmt")

func main( ) {
	fruits := [3]string{"orange", "apple", "banana"}

	for index, _ := range fruits {
		fmt.Printf("%v\n", index)
	}
}
