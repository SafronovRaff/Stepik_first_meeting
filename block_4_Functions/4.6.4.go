/*

Напишите функцию double_m(), которая должна принимать на вход два целых числа a и b и возвращать сумму квадратов чисел от a до b (включительно).

Sample Input:

2 6
Sample Output:

90

 */

func double_m(a, b int) int {
Sum := 0
if a < b {
for i := a; i <= b; i++ {
Sum += i * i
}
} else {
for i := b; i <= a; i++ {
Sum += i * i
}
}
return Sum
}
