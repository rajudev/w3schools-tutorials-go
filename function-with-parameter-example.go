package main
import ("fmt")

func familyName(fname string, age int) {
	fmt.Println("Hello", fname, "Wood", age, "years")
}

func main(){
	familyName("Jos", 23)
	familyName("Hazel", 34)
	familyName("Mark", 25)
}
