// Problem : http://codeforces.com/contest/787/problem/A
// Tages : Implementation , Math
/*
Forgetting "/n" in the line 11 made my solution get a wrong answer
because the online-judge give the input in two seprated lines
*/
package main 
import ."fmt"

func main(){
    var a , b , c , d int 
    Scanf("%d%d\n" , &a , &b)
    Scanf("%d%d" , &c , &d)
    mapper := make([]int , 1e6+5)
    for i:=b; i<= 1e6; i+=a{
        mapper[i] += 1
    }
    ans := -1
    for i:=d; i<= 1e6; i+=c{
        mapper[i] += 1
        if mapper[i] == 2{
            ans = i
            break
        }
    }
    Printf("%d" , ans) 
}