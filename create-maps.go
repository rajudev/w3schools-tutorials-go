package main
import ("fmt")

func main() {
	var a = map[string]string{"brand":"Ford", "model": "mustang", "year":"1964"}
	b := map[string]int{"berlin":1, "munich":2, "singapore":3, "Schengen":4}

	fmt.Printf("a\t%v\t\n", a)
	fmt.Printf("b\t%v\t\n", b)
}
