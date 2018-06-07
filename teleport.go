package main

import (
	"fmt"
	"os"
	"sync"
)



type path struct {
	time int			//время в пути
	route []int			//маршрут - список посещенных точек
	here int			//точка, в которой мы находимся в данный момент
}


//Метод проверяет посещался ли уже нами пункт direct
func (p path) isVisited (direct int) bool {
	for _,visited := range p.route {
		if direct == visited {
			return true
		}
	}
	return false
}


//Метод добавляет точку в маршрут,
// обозначет эту точку как нынешнюю
// и добавляет к времени пути время обслуживания в этой точке
func (p path) addPoint (direct int, d *data) (path) {
	p.here = direct
	p.time += d.stations[direct].service
	p.route = append (p.route, direct)
	return p
}




type point struct {
	service int			//время обслуживания в данной точке
	teleport []int		//список точек, в которые можно телепортироваться из данной точки
}


type data struct {
	stations [] *point		//список точек
	minTime int				//минимальное время уже известных маршрутов
	wg sync.WaitGroup
}


//Функция определяет нужно ли идти дальше или мы уже прибыли в пункт назначения,
// и в соответствии с этим либо отправляет в канал время пути,
// либо продолжает путь.
//
//(Оптимизирована тем, что путь не продолжается в точки, где мы уже были ранее,
// а также путь не продолжается, если время в пути уже стало больше наименьшего из известных)
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

//Определяет наименьшее время маршрутов "на финише"
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

	if d.minTime != 0 {
		fmt.Println(d.minTime - d.stations[N-1].service)
	}else{
		fmt.Println(-1)
	}

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
