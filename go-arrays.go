package main

import ("fmt")

func main() {
	var arr1 = [...]int{1, 2, 3}
	var arr2 = [...]int{5,6,7,8}
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(arr2)
	fmt.Println(len(arr2))
	fmt.Println(cars)
	fmt.Println(len(cars))
}
