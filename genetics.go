package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	var str string
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	str = myscanner.Text()

	data := []rune(str)
	aRune := []rune("@")[0]

	first := -1
	last := -1
	for i := 0; i < len(data); i++ {
		if data[i] == aRune && first == -1 {
			first = i
		}
		j := len(data) - i - 1
		if data[j] == aRune && last == -1 {
			last = j
		}
		if first != -1 && last != -1 {
			break
		}
	}
	sub := data[first+1:last]

	for i := 0; i < len(sub)/2; i++ {
		j:=len(sub)-i-1
		sub[i], sub[j] = sub[j], sub[i]
	}

	fmt.Println(string(data))
}
