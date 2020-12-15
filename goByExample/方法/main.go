package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func(r *rect) area() int{
	return r.width * r.height
}

func (r rect) perim() int{
	r.width = 100
	return 2 * r.width * r.height

}

func main(){
	r := rect{10,10}

	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r

	fmt.Println(rp.width)
}