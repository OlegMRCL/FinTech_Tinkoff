package main

import (
	"fmt"
	"os"
)

type skillTable [][] int
func (t *skillTable) giveJob (s []int, taskNum int) {
	for worker, skill := range *t {
		if (skill[taskNum] == 1)&&(!busy(worker, s)) {
			go t.giveJob(append(s, worker), taskNum+1)
		}
	}
}

func busy(w int, s []int) bool {
	for _,v:=range s {
		if v == w {
			return true
		}
	}
	return false
}

func main () {
	var N, M int
	fmt.Fscan(os.Stdin, &N, &M)


	t:=new(skillTable)

	for i:=0; i<N; i++ {
		tasks:=make([]int, N)
		for j:=0; j<M; j++ {
			fmt.Fscan(os.Stdin, &tasks[j])
		}
		*t[i]=tasks
	}

	taskNum:=0
	s:=make([]int, N)

	go

}
