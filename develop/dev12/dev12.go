// 12.	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func main() {
	strings := [5]string{"cat", "cat", "dog", "cat", "tree"}
	result := make(map[string]struct{})

	for _, v := range strings {
		result[v] = struct{}{}
	}
	for k, _ := range result {
		fmt.Println(k)
	}
}
