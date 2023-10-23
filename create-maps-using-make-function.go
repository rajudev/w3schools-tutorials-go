package main
import ("fmt")

func main(){
	var a=make(map[string]string) // a is empty here

	a["brand"] = "Ford"
	a["Model"] = "Mustang"
	a["year"] = "1964"

	// a is no longer empty

	b := make(map[string]int)

	b["Berlin"] = 1
	b["Munich"] = 2
	b["Singapore"] = 3
	b["Schengen"] = 4

	fmt.Printf("a\t%v\t\n", a)
	fmt.Printf("b\t%v\t\n", b)
}
