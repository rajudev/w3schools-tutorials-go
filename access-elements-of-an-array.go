package main
import ("fmt")

func main() {
	prices := [3]int{30,60,90}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
	fmt.Println("Changing value of index 1")

	prices[1] = 50
	fmt.Println(prices)
}
