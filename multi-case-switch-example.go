package main
import ("fmt")

func main() {
	day := 9

	switch day {
	case 1,2,3,4,5:
		fmt.Println("Day is a weekday")
	case 6,7:
		fmt.Println("Day is a weekend")
	default:
		fmt.Println("Invalid Day Number")
	}
}
