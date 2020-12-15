package main

import "fmt"

func main(){
	messager := make(chan string)

	go func() {
		messager <- "ping"
	}()
	a := <- messager
	fmt.Println(a)
	fmt.Println(messager)
}
