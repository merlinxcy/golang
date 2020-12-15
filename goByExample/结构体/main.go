package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){
	fmt.Println(person{"123",2})
	fmt.Println(person{name: "123"})
	fmt.Println(person{name: "123",age: 1})
	fmt.Println(&person{name: "123",age:11})

	s := person{name:"123",age: 111}
	fmt.Println(s.name)

	pointer := &s
	pointer.age = 11111111

	fmt.Println(s.age)
	fmt.Println(pointer.age)
}