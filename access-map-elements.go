package main
import ("fmt")

func main() {
	a := make(map[string]string)

	a["Brand"] = "Ford"
	a["Model"] = "Mustang"
	a["year"] = "1965"

	fmt.Printf(a["Brand"])
}
