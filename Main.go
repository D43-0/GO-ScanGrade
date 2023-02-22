//
// Create Date Feb 22 2023
// If method

package main

import	(
		"fmt"
		)
		
func main() {
		fmt.Println("Grade your value")
		fmt.Println("")
		var value int
		
		fmt.Println("Input Value? 1-100")
		fmt.Scan(&value)
		
	if value < 25 {
		fmt.Println("Grade D");
		
	} else if value < 51 {
		fmt.Println("Grade C");
		
	} else if value < 76 {
		fmt.Println("Grade B");
		
	} else if value <= 100 {
		fmt.Println("Grade A");
		
	} else {
		fmt.Println("Grade Error");
	}
}