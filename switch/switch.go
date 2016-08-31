package main 

import "fmt"
import "time"

func main() {
	
	fmt.Println("Today is", time.Now())

	fmt.Println("Today is", time.Now().Weekday())

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("Today is Weekend")
		default :
			fmt.Println("Today is Weekday") 

	}
}