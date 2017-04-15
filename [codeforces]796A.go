// Problem : http://codeforces.com/contest/796/problem/A
// Tages : Implementation , brute-force
package main
import (
	"fmt"
	"math"
)

func main(){
	var n,m,K,len,pos int32
	fmt.Scanf("%d%d%d\n",&n,&m,&K)
	m--
	len,pos= 123456789,123456789
	var i int32
	for i=0;i<n;i++{
		var price int32
		fmt.Scanf("%d",&price)
		if price>0&& price<=K{
			if int32(math.Abs(float64(m-i)))< len{
				len,pos=int32(math.Abs(float64(m-i))),i
			}
		}
	}
	fmt.Printf("%d\n",int32(math.Abs(float64(pos-m)))*10)
}
