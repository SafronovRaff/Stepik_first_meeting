/*
Вам нужно написать функцию one_or_two(),  которая принимает два целых числа и строку.
Строка может иметь значения one, two или любой другой текст.
Возвращать из функции вам нужно два значения. Если строка равна one, нужно вернуть первое число и саму строку.
Если строка равна two, нужно вернуть второе число и саму строку. Если строка другая - нужно вернуть 0 и саму строку.


Sample Input:

2 5 two
Sample Output:
5 two

 */
func one_or_two(a, b int, c string) (int, string) {

switch c {
	case "one":
		return a, c
	case "two":
		return b, c
	default:
		return 0, c
	}

}

