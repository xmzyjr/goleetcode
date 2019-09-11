package main

import "fmt"

func kthGrammar(N int, K int) int {
	if N == 1 && K == 1 {
		return 0
	}
	i := 1 << uint(N-2)
	if K > i {
		return 1 - kthGrammar(N-1, K-i)
	} else {
		return kthGrammar(N-1, K)
	}
}

func main() {

	fmt.Println(kthGrammar(3, 3))
}
