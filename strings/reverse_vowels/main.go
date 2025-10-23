package main

import (
	"fmt"
)


var vowels=map[byte]struct{}{
    'a':{},
    'e':{},
    'i':{},
    'o':{},
    'u':{},
    'A':{},
    'E':{},
    'I':{},
    'O':{},
    'U':{},
}

func reverseVowels(s string) string {
    var l,r int
    r=len(s)-1
    result:=[]byte(s)
    for ;l<r;{
        if _,ok:=vowels[result[l]];!ok{
            l++
            continue
        }
        if _,ok:=vowels[result[r]];!ok{
            r--
            continue
        }
        result[l],result[r]=result[r],result[l]
        l++
        r--
    }
    return string(result)
}

func main() {
    fmt.Println(reverseVowels("teste")
}