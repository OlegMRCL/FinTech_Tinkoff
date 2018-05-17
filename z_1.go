package main

import (
	"fmt"
	"os"
)

func main() {
	var K,M int
	fmt.Fscan(os.Stdin, &K, &M)

	if K!=M {
		fmt.Println(K+M-1)
	}else{
		fmt.Println(0)
	}
}
