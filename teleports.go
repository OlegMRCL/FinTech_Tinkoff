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

func (p path) addPoint (direct int, d *data) (path) {
	p.here = direct
	p.time += d.stations[direct].service
	p.route = append (p.route, direct)
	return p
}




type point struct {
	service int
	teleport []int
}


type data struct {
	stations [] *point
	minTime int
	wg sync.WaitGroup
}

func (d *data) jump (p path, ch chan int) {
	defer d.wg.Done()
	if p.here == 0 {
		ch <- p.time
	}else{
		s := d.stations[p.here]
		for _,direct := range s.teleport {
			if !p.isVisited(direct) {
				newPath := p.addPoint (direct, d)
				if (newPath.time < d.minTime)||(d.minTime == 0){
					d.wg.Add(1)
					go d.jump(newPath, ch)
				}
			}
		}
	}
}

func(d *data) finish (ch chan int) {
	for time := range ch {
		if (time < d.minTime) || (d.minTime == 0) {
			d.minTime = time
		}
	}
}

func (d *data) calculate(N int) {
	ch := make (chan int, 1)

	var newPath path
	newPath = newPath.addPoint(N-1, d)
	d.wg.Add(1)
	go d.jump(newPath, ch)

	go d.finish(ch)
	d.wg.Wait()
	close(ch)

	fmt.Println(d.minTime-d.stations[0].service)
}







func main() {

	var N int
	d := new(data)
	fmt.Fscan(os.Stdin, &N)
	d.stations = make ([] *point, N)
	for i:=0; i<N; i++ {
		s := new(point)
		fmt.Fscan(os.Stdin, &s.service)
		d.stations[i] = s
	}
	var M int
	fmt.Fscan(os.Stdin, &M)
	for i:=0; i<M; i++ {
		var a, b int
		fmt.Fscan(os.Stdin, &a, &b)
		p := d.stations[a-1]
		p.teleport = append (p.teleport, b-1)
		p = d.stations[b-1]
		p.teleport = append (p.teleport, a-1)
	}


	d.calculate(N)


}
