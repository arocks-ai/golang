package main

import (
	"fmt"
	"time"
	"reflect"
)


func main(){

	// Very basic switch usage
	fmt.Println("== Switch usage with int type")
	i := 2
	switch i {
	case 2:
		fmt.Println("value is 2")
	case 1:
		fmt.Println("value is 1")
	default:
		fmt.Println("unknown value")
	}

	
	// Switch example with time and string
	fmt.Println("\n== Switch example with time and string ...")
	today := time.Now().Weekday()
	fmt.Println("today = ", today, "Type of today is = ", reflect.TypeOf(today))	// type is time.Weekday
	
	switch today.String() {			// Convert to String for compare pupose
	case "Sunday", "Saturday":		// Multiple string compare	
		fmt.Println("Weeekend")
	default:
		fmt.Println("Weekday")
	}

	
	// Switch example with better version of time method usage ...
	fmt.Println("\n== Switch example with better version of time method usage ...")
	
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Weeekend")
	default:
		fmt.Println("Weekday")
	}


	// Switch example with better version of time method usage ...
	fmt.Println("\n== Switch example with no expression ...")
	t := time.Now()
	fmt.Println("time = ", t)

	switch {	// No expression
	case t.Hour()  < 12:
		fmt.Println("AM")
	default:
		fmt.Println("PM")
	}

	switch t.Month() {	// No expression
	case 1:
		fmt.Println("Jan")
	case 4:
		fmt.Println("Apr")		
	default:
		fmt.Println("Some other month")
	}

}