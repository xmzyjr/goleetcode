package main

import "fmt"

func countCharacters(words []string, chars string) int {
	var hash [26]int
	for _, v := range chars {
		hash[v-'a']++
	}
	count := 0
label:
	for _, v := range words {
		var temp = hash
		for _, v1 := range v {
			if temp[v1-'a'] == 0 {
				continue label
			} else {
				temp[v1-'a']--
			}
		}
		count += len(v)
	}
	return count
}

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
}
