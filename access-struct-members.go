package main
import ("fmt")

type Person struct {
	name string
	age int
	job string
	salary int
}

func main() {
	var person1 Person
	var person2 Person

	person1.name = "Bob"
	person1.age = 23
	person1.job = "Engineer"
	person1.salary = 40000

	person2.name = "Alice"
	person2.age = 25
	person2.job = "Engineer"
	person2.salary = 34000

	printPerson(person1)
	printPerson(person2)
}

func printPerson(pers Person){
	fmt.Println(pers.name)
	fmt.Println(pers.age)
	fmt.Println(pers.job)
	fmt.Println(pers.salary)
	
}
