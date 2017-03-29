// Problem : https://www.hackerrank.com/contests/w30/challenges/candy-replenishing-robot
// Type : Implementation
package main
import "fmt"



func main(){
	var n , t int
	fmt.Scanf("%d%d" , &n , &t)
	all := n 
	ans := 0
	for i := 0 ; i < t ; i++{
		var _var int
		fmt.Scanf("%d" , &_var)
		n -= _var
		if( n < 5 && i != t - 1){
			ans += all - n;
			n = all
		}
	
	//fmt.Printf("%d\n" , ans)
	}	
	fmt.Println(ans)


}