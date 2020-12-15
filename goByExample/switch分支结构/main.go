package main

import "fmt"
import "time"

func main(){
	i := 2
	fmt.Print("write", i , " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("it is the weeend")
	default:
		fmt.Println("it is a weekend")
	}
	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it's a weekday")
	}

	whatAmi := func(i interface{}){
		switch t := i.(type) {
		case bool:
			fmt.Println("l am a bool")
		case int:
			fmt.Println("l am a int")
		default:
			fmt.Printf("don't know type %T\n", t)
		}
	}
	whatAmi(true)
	whatAmi(1)
	whatAmi("hey")
}
