// In tribute to my mum, Angel Giris
// Problem :https://www.hackerrank.com/contests/w32/challenges/duplication
// Tages : Implementation , strings
/*
*/
package main 
import "fmt"
//import "strings"

func fib(s string) string{
    ret := ""
    for i:=0 ; i < len(s) ; i++{
      if s[i] == '1'{
        ret += "0"
      }else{
        ret += "1"
      }
    }
    return ret 
}
func main(){
  s := "0"
  for len(s) <= 1000 {
    s += fib(s)
  }
  var tc int 
  fmt.Scan(&tc)
  for i:=0 ; i < tc ; i++{
    var tcase int 
    fmt.Scan(&tcase)
    fmt.Printf("%c\n" , s[tcase])
  }
}
