package main

import "fmt"


func doA(){
	//panic("error")
	fmt.Println("do A")
}


func doB(){
	fmt.Println("do B")
}


func main(){
	doA()
	fmt.Println("ok")
	defer doB()
	fmt.Println("final?")
}
