package main
import ("fmt")

func main() {
	var a = map[string]string {"brand":"Ford", "Model":"Mustang", "year":"1970"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1981"

	fmt.Println(a)
	fmt.Println(b)
	
}
