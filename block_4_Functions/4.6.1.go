/*
Напишите функцию max() которая должна принимать на вход два целых числа и возвращать большее из них.

Sample Input:

10 4
Sample Output:

10
*/

func max(a, b int) int {

if a > b {
return a
} else if b > a {
return b

}
return 0
}