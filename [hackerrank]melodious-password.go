package main
import "fmt"
func createPW(length int, pre string, vowelsPrev bool) {
    if length > 0 {
        if vowelsPrev {
            for _,consonant := range consonants {
                createPW(length-1, pre+consonant,false)
            }
        } else {
            for _,vowel := range vowels {
                createPW(length-1, pre+vowel,true)
            }
        }
    } else {
        fmt.Printf("%s\n",pre)
    }
}

var vowels, consonants []string
func main() {
    var n int
    fmt.Scanf("%d",&n)
    vowels = []string{"a","e","i","o","u"}
    consonants = []string{"b","c","d","f","g","h","j","k","l","m","n","p","q","r","s","t","v","w","x","z"}
    createPW(n,"",false)
    createPW(n,"",true)
}