package main

import (
	"fmt"
	"strings"

	"example.com/roman2int/roman2int"
)

func main() {
	var romanNo string
	for {
		fmt.Println("Enter Roman String in Capital letter:")
		fmt.Scan(&romanNo)
		romanNo = strings.ToUpper(romanNo)
		fmt.Printf("%d\n\n", roman2int.Roman2Int(romanNo))
		fmt.Printf("%d\n\n", roman2int.Roman2Intv2(romanNo))
		fmt.Printf("%d\n\n", roman2int.Roman2IntOptimize(romanNo))
	}
}
