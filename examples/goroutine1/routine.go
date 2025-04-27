/* Very basic example of goroutine

demo - showing how to invoke go routine
main() will return before go routine finishes as there is no wait for go routine to finish

dbConnect() method simiulates connecting to database by sleeping for random seconds.
We are calling dbConnect() method multiple times using go routine.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"db1", "db2", "db3", "db4", "db5"} // Global data

func main() {
	for i := 0; i < 5; i++ {
		go dbConnect(i) // Calling go routine
	}

	fmt.Println("= Inside main() method ...")
	// for j := range len(dbData) {
	// 	fmt.Print(dbData[j] + " ")
	// }
	fmt.Println()
}

// Function to simulate some delay 
func dbConnect(id int) {
	//fmt.Println("Connecting to database", connID)
	randomSec := rand.Intn(5) + 1                      // Random number between 1 to 5
	time.Sleep(time.Duration(randomSec) * time.Second) // Sleep for random seconds
	fmt.Println("== Connected to database", dbData[id]) // Unreachable code as main() func exits before the delay   	
}
