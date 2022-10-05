/*
Результаты матчей
Вы создаете программу для анализа результатов спортивных матчей и подсчета очков заданной команды.
Результаты матчей хранятся в массиве results.
Каждый матч имеет один из следующих результатов:
"w" - выиграл
"l" - проиграл
"d" - ничья

Победа добавляет три очка, ничья - одно очко, а проигранный матч не добавляет очков.

Ваша программа должна принять на вход результат последнего матча и добавить его в массив результатов. После этого необходимо вычислить и вывести общее количество очков, набранных командой по результатам матчей из массива results.

Sample Input:

w
Sample Output:

22
*/
package main

import "fmt"

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	num := 0
	var res string
	fmt.Scanln(&res)
	switch res {
	case "w":
		results = append(results, "w")
	case "l":
		results = append(results, "l")
	case "d":
		results = append(results, "d")
	}
	info := map[string]int{"w": 3,
		"d": 1,
		"l": 0}
	for _, v := range results {
		num += info[v]
	}
	fmt.Println(num)
}
