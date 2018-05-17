package main

import (
	"fmt"
	"os"
	"sync"
)



type path struct {
	time int
	route []int
	here int
}
func (p path) isVisited (direct int) bool {
	for _,visited := range p.route {
		if direct == visited {
			return true
		}
	}
	return false
}
func (p path) addPoint (direct int) (newPath path) {
	newPath = p
	newPath.here = direct
	newPath.time += stations[direct].service
	newPath.route = append (newPath.route, direct)
	return
}




type point struct {
	service int
	teleport []int
}
var stations [] *point





var minTime int

func jump (p path, ch chan int) {
	defer wg.Done()
	if p.here == 0 {
		ch <- p.time
	}else{
		s := stations[p.here]
		for _,direct := range s.teleport {
			if !p.isVisited(direct) {
				newPath := p.addPoint (direct)
				if (newPath.time < minTime)||(minTime == 0){
					wg.Add(1)
					go jump(newPath, ch)
				}
			}
		}
	}
}

func finish (ch chan int) {
	time, ok := <-ch
	if ok {

		if (time < minTime)||(minTime == 0) {
			minTime = time
		}

		go finish(ch)
	}


}




var wg sync.WaitGroup


func main() {

	var N int
	fmt.Fscan(os.Stdin, &N)
	stations = make ([] *point, N)
	for i:=0; i<N; i++ {
		s := new(point)
		fmt.Fscan(os.Stdin, &s.service)
		stations[i] = s
	}
	var M int
	fmt.Fscan(os.Stdin, &M)
	for i:=0; i<M; i++ {
		var a, b int
		fmt.Fscan(os.Stdin, &a, &b)
		p := stations[a-1]
		p.teleport = append (p.teleport, b-1)
		p = stations[b-1]
		p.teleport = append (p.teleport, a-1)
	}

	for i:=0; i<N; i++ {
		m := *stations[i]
		fmt.Println(m)
	}




	ch := make (chan int, 1)


	var newPath path
	newPath = newPath.addPoint(N-1)
	fmt.Println(newPath)
	wg.Add(1)
	jump(newPath, ch)

	finish(ch)
	wg.Wait()
	close(ch)

	fmt.Println(minTime-stations[0].service)
}
