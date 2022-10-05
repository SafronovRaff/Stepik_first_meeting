/*
Легкие задания остались позади. Пора решить что-то посложнее. Итак...

Вы создаете робота, который может говорить цифры.
Ваш робот должен принимать на вход 3 целых числа (каждый с новой строки) в диапазоне от 0 до 10 (включительно)
и выдавать соответствующий текст на русском языке.
Смотрите пример в Sample Input и Sample Output. Не забывайте, на вход могут подаваться разные числа,
а ваше решение должно работать с любыми из них, не только с теми, которые в примере.
0 выводим как Ноль

Sample Input:
3
1
5
Sample Output:
Три
Один
Пять
*/
package main

import "fmt"

func main() {
	var a int
	for i := 0; i < 3; i++ {
		fmt.Scan(&a)
		switch a {
		case 0:
			fmt.Println("Ноль")
		case 1:
			fmt.Println("Один")
		case 2:
			fmt.Println("Два")
		case 3:
			fmt.Println("Три")
		case 4:
			fmt.Println("Четыре")
		case 5:
			fmt.Println("Пять")
		case 6:
			fmt.Println("Шесть")
		case 7:
			fmt.Println("Семь")
		case 8:
			fmt.Println("Восемь")
		case 9:
			fmt.Println("Девять")
		case 10:
			fmt.Println("Десять")
		}
	}
}
