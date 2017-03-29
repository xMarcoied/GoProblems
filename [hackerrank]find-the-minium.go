// Problem : https://www.hackerrank.com/contests/w30/challenges/find-the-minimum-number
// Recursion
package main
import ."fmt"


func solve(n int){
	if n == 1{
		Printf("int")
		return
	}
	Printf("min(int, ")
	solve(n-1)
	Printf(")")
}

func main(){
	var n int
	Scanf("%d" , &n)
	solve(n)
}