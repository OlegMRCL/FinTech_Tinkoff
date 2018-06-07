package main

import (
	"fmt"
	"os"
)

type data struct {
	heads map [string] []string
}

func (d *data) printHeads (s []string, lvl int) {
	var ns []string
	for _,head:=range s {

		fmt.Println(head, lvl)

		for _,name:=range d.heads[head] {
			ns = append (ns, name)
		}
		if len(ns)>0 {
			d.printHeads(ns, lvl+1)
		}

	}
}

func main () {

	var N int
	fmt.Fscan(os.Stdin, &N)

	d:=new(data)

	d.heads=make(map[string] []string)

	for i:=1; i<N; i++ {
		var name, head string
		fmt.Fscan(os.Stdin, &name, &head)
		d.heads[head] = append (d.heads[head], name)
	}

	s:=make([]string, 1)
	s[0] = "X"

	d.printHeads(s, 0)

}
