package main
import ("fmt")

func main(){
	var a = map[string]string{
		"brand":"Ford",
		"model":"mustang",
		"year":"1965",
		"day":""}

	val1, ok1 := a["brand"] //checking for existing key and its value
	val2, ok2 := a["color"] //check for non-existing key and its value
	val3, ok3 := a["day"] //check for existing key and its value
	_, ok4 := a["model"] //only checking for existing key and not its value
	fmt.Println(val1,ok1)
	fmt.Println(val2,ok2)
	fmt.Println(val3,ok3)
	fmt.Println(ok4)
	
}
