package main

import "fmt"

func maddd(numsss ...int){
	fmt.Println(numsss)
}

func main(){
	test := []int{1,2,3,5,6}
	maddd(test...)
}