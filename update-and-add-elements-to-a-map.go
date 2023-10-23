package main
import ("fmt")

func main() {
	var a = make(map[string]string)

	a["Brand"] = "Ford"
	a["Model"] = "Mustang"
	a["year"] = "1964"

	fmt.Println(a)

	a["Color"] = "Black"

	fmt.Println(a)


	
}
