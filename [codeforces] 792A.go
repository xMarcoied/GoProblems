// Problem : http://codeforces.com/contest/792/problem/A
// Tags : Sorting
/*
Notes :
*/
package main 
import ."fmt"
import "sort"
import "math"
func main(){
  var n int
  Scanf("%d\n" , &n)
  arr := make([]int , n)
  for i:=0 ; i < n ; i++{
    Scanf("%d" , &arr[i])
  }
  sort.Ints(arr)
  var ans , cnt float64 
  ans = 1e9+10
  for i:= 1 ; i < n ; i++{
    reg := (arr[i] - arr[i-1])
    dist := math.Abs(float64(reg))
    if dist == ans {
      cnt += 1
    }else if dist < ans {
      ans = dist 
      cnt = 1 
    }
  }
  Printf("%.0f %.0f" , ans , cnt)
}