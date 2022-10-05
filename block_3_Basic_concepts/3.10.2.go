/*
package main

import "fmt"

	func main() {
	    height := 195
	    __ height > 185 {
	        fmt.Println("Да")
	    } ____ {
	        fmt.Println("Нет")
	    _
	}
*/
package main

import "fmt"

func main() {
	height := 195
	if height > 185 {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}
}
