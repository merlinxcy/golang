package main

import (
	"fmt"
	"github.com/sensepost/godoh/utils"
)
import "time"
func producer(message chan string){
	fmt.Println(message)
	for {
		time.Sleep(time.Second*5)
		message <- "test:" + utils.RandomString(3)
		fmt.Println("[+]producer")
	}
}

func comsummer(message chan string){
	fmt.Println(message)
	for{
		time.Sleep(time.Second)
		fmt.Println(<-message)//阻塞的
		fmt.Println("[-]comsummer")
	}
}


func main(){
	message := make(chan string,2)

	go comsummer(message)
	go producer(message)

	time.Sleep(time.Hour)
}
