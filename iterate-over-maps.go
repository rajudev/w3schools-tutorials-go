package main
import ("fmt")

func main() {
	var a = map[string]int{"Singapore":1, "Berlin":2, "Shengen":3, "Kigali": 4}

	for k, v := range a {
		fmt.Printf("%v\t%v\n", k, v)
	}
}
