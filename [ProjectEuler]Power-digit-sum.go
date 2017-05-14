// Problem : https://www.hackerrank.com/contests/projecteuler/challenges/euler016
// Tages : Implementation , Math , casting
/*
Making the power function taking Int
*/
package main 
import (
	"fmt"
	"math"
)
func solve(in float64) int{
  ret := 0 
  n := int(in)
  for n > 0 {
    d := n % 10
    n /= 10 
    ret += d
  }
  return ret
}
func main(){
  var tc int 
  var x float64
  fmt.Scanf("%d" , &tc)
  for tcase :=0 ; tcase < tc ; tcase ++{
      fmt.Scan(&x)
      x = math.Pow(2 , x)
      fmt.Println(solve(x))
  }
}
