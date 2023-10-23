package main
import ("fmt")

func main() {
	day := 5

	switch day {
	case 1: fmt.Println("Day is Monday")
	case 2: fmt.Println("Day is Tuesday")
	case 3: fmt.Println("Day is Wednesday")
	case 4: fmt.Println("Day is Thursday")
	case 5: fmt.Println("Day is Friday")

	default:fmt.Println("Not a Weekday")
	}
}
