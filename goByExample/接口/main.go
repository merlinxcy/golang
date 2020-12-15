package main

import "fmt"

type geometry interface {
	func1() int
	func2() int
}

type rect struct {
	aaa int
	bbb int
}

func(r *rect)func1()int{
	return r.aaa + r.bbb
}

func(r *rect)func2()int{
	return  1
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.func1())
}


func main(){
	r := rect{1,2}
	measure(&r)
}