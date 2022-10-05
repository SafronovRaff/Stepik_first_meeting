/*
Произведение всех
На вход подается целое число. Вам необходимо вывести произведение всех чисел от 1 до данного числа (включительно).

Например, на вход подается число 5, вам нужно найти - 1*2*3*4*5 = 120

Sample Input:

4
Sample Output:

24
*/
package main

import "fmt"

func main() {
	var a int
	var sum = 1
	fmt.Scanln(&a)

	for i := 1; i <= a; i++ {
		sum *= i
	}
	fmt.Println(sum)
}
