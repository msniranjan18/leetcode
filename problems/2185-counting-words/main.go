package main

import (
	"fmt"
)

func main() {
	fmt.Println(prefixCount([]string{"pay","attention","practice","attend"}, "at"))
}

func prefixCount(words []string, pref string) int {
    count:=0
    for _,word := range words {
        if len(word)<len(pref) {
            continue
        }
        i:=0
        for ;i<len(pref); i++ {
            if pref[i] != word[i] {
                break
            }
        }
        if i==len(pref) {
            count++
        }
    }
    return count
}