// In tribute to my mum, Angel Girl 
/*
Problem : https://www.hackerrank.com/contests/w32/challenges/fight-the-monsters
Tags : Sorting , Math 
*/
package main 
import (
  "fmt"
  "sort"
  "math"
)

func main(){
  var n , h , t , ans int 
  fmt.Scan(&n , &h , &t)
  arr := make([]int , n)
  for i:= 0 ; i < n ; i++{
    fmt.Scan(&arr[i])
  }
  sort.Ints(arr)
  for i:=0 ; i < n ; i++{
    can := math.Ceil(float64(arr[i]) / float64(h))
    //fmt.Println(can)
    if(t >= int(can)){
      ans += 1 
      t -= int(can)
    }
  }
  fmt.Println(ans)
}
