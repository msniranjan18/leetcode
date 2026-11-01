package main

import (
  "fmt"
)

func main() {
  fmt.Println(partitionString("abacaba"))
}

func partitionString(s string) int {
    return sol1(s)
    return sol2(s)
}

func sol2(s string) int {
    checker := [26]int{} // fixed array for 26 alphabets 
    count:=0
    for i:=0; i<len(s); i++ {
        if checker[s[i]-'a'] > 0 {
            count++
            checker = [26]int{} // reset the map for the next substring
        }
        checker[s[i]-'a']++
    }
    return count+1
}

func sol1(s string) int {
    checker := make(map[byte]struct{})
    count:=0
    for i:=0; i<len(s); i++ {
        if _,ok:=checker[s[i]]; ok {
            count++
            clear(checker) // reset the map for the next substring
        }
        checker[s[i]] = struct{}{}
    }
    return count+1
}
