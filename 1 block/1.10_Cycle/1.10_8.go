package main

import "fmt"

/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

# Входные данные

Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.

# Выходные данные

Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.

Внимание! Если вам сложно решить эту задачу, пропустите и изучайте материал дальше, затем вернетесь к этой задаче позже.

Sample Input:

564 8954
Sample Output:

5 4
*/
func main() {

	var x, y string
	var i, j int

	fmt.Scan(&x, &y)

	for i = 0; i < len(x); i++ {
		for j = 0; j < len(y); j++ {
			if string(x[i]) != string(y[j]) {
				continue
			}
			fmt.Print(string(x[i]), " ")
		}

	}

}