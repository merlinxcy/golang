package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i:=0; i<3; i++{
		fmt.Println(from,":",i)
	}
}

func  main()  {
	f("direct")

	go f("test")
	go func(msg string) {
		time.Sleep(12)
		fmt.Println(msg)
	}("ttttghygycvgsdcvf")

	time.Sleep(time.Second)
	fmt.Println("done")
}