/*
Богач или бедняк?
На вход подается целое число, сумма денег, которые у вас есть.
Ваша задача - вывести марку телефона, которую вы можете себе позволить купить.

Если сумма больше 1000 - вывести Apple
Если сумма от 500 до 1000 (включительно) - вывести Samsung
Если сумма меньше 500 - вывести Nokia с фонариком

Sample Input:

1100
Sample Output:

Apple
*/
package main

import "fmt"

func main() {
	var sum int
	a, b := 1000, 500
	fmt.Scanln(&sum)

	switch {
	case sum > a:
		fmt.Println("Apple")
	case sum <= a && sum >= b:
		fmt.Println("Samsung")
	case sum < b:
		fmt.Println("Nokia с фонариком")
	}
}
