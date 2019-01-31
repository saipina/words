// // hello/count.go
// package main

// import (
// 	"fmt"
// 	"words"
// )

// func main() {
// 	s := "if it walk like a duck..."
// 	w := words.WordCount(s)
// 	fmt.Println(s)
// }
package words

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	fmt.Println("kad")
	words := strings.Fields(s) // HL
	r := map[string]int{}
	for _, w := range words {
		r[w] = r[w] + 1 // HL
	}
	return r
}
