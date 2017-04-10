// Problem :https://code.google.com/codejam/contest/3264486/dashboard#s=p1
// Tags : Implementation , Maths
package main
import "fmt"
func valid(n int) bool{
  prev := n % 10
  n /= 10
  for n > 0{
    d := n % 10
    n /= 10 
    if d > prev {
      return false
    }
    prev = d 
    
  }
  return true
}
func main() {
    var tc int 
    fmt.Scanf("%d\n" , &tc)
    for testcase:=1 ; testcase <= tc ; testcase ++{
        fmt.Printf("Case #%d: " , testcase)
        var x int 
        ans := 0
        fmt.Scanf("%d\n" , &x)
        for i:= x ; i >= 1 ; i--{
            if valid(i){
              ans = i 
              break
            } 
        }
        fmt.Printf("%d\n" , ans)
        
    }
}