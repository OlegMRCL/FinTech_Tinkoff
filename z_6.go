package main

import (
	"fmt"
	"os"
	"math"
)


func inHead (x,y float64) bool {
	return (math.Pow(x,2)/2 + math.Pow(y, 2)) < 1
}

func belowBrows (x,y float64) bool {
	return (0.5 * math.Abs(x) + 0.5) > y
}

func outOfEyes (x,y float64) bool {
	return ((math.Pow((x - 0.5), 2) + math.Pow(y, 2)) > 0.3)&&((math.Pow((x + 0.5), 2) + math.Pow(y, 2)) > 0.3)
}

func success (x,y float64) bool {
	return inHead(x,y)&&belowBrows(x,y)&&outOfEyes(x,y)
}

func main () {
	var x,y float64
	fmt.Fscan(os.Stdin, &x, &y)



	if success(x,y) {
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
		}

}