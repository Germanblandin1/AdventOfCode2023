package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

var basura string
var t int
var n, m int
var INF int = 1000000000
var actual int

var matriz [][]rune

type Pair struct {
	i, j, can, ant_dir int
}

type Item struct {
	value    Pair
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(*pq) == 0
}

var movi = []int{0, 0, 1, -1}
var movj = []int{1, -1, 0, 0}

var dist [][][][]int

var max_dist_dir = 10
var min_dist_dir = 4

func dijkstra(i, j, dir int) int {
	dist = make([][][][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([][][]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = make([][]int, max_dist_dir+1)
			for k := 0; k < max_dist_dir+1; k++ {
				dist[i][j][k] = make([]int, 4)
				for l := 0; l < 4; l++ {
					dist[i][j][k][l] = INF
				}
			}
		}
	}

	pq := PriorityQueue{}
	heap.Init(&pq)

	dist[i][j][0][dir] = 0
	item := &Item{
		value:    Pair{i, j, 0, dir},
		priority: 0,
	}
	heap.Push(&pq, item)

	for !pq.IsEmpty() {
		item := heap.Pop(&pq).(*Item)
		u := item.value
		d := item.priority
		//fmt.Println(u)
		if u.i == n-1 && u.j == m-1 && u.can >= min_dist_dir {
			return d
		}

		for k := 0; k < 4; k++ {
			ii := u.i + movi[k]
			jj := u.j + movj[k]
			if ii >= 0 && ii < n && jj >= 0 && jj < m {

				nextcan := 1
				if u.ant_dir == k {
					nextcan = u.can + 1
				}

				if nextcan > max_dist_dir {
					continue
				}
				if u.ant_dir == 0 && k == 1 || u.ant_dir == 1 && k == 0 || u.ant_dir == 2 && k == 3 || u.ant_dir == 3 && k == 2 {
					continue
				}
				if u.ant_dir == 0 && (k == 2 || k == 3) || u.ant_dir == 1 && (k == 2 || k == 3) || u.ant_dir == 2 && (k == 0 || k == 1) || u.ant_dir == 3 && (k == 0 || k == 1) {
					if u.can < min_dist_dir {
						continue
					}
				}

				val, _ := strconv.Atoi(string(matriz[ii][jj]))
				if dist[u.i][u.j][u.can][u.ant_dir]+val < dist[ii][jj][nextcan][k] {
					dist[ii][jj][nextcan][k] = dist[u.i][u.j][u.can][u.ant_dir] + val
					heap.Push(&pq, &Item{
						value:    Pair{ii, jj, nextcan, k},
						priority: dist[ii][jj][nextcan][k],
					})
				}
			}
		}
	}
	return INF

}

func main() {
	fmt.Scan(&n, &m)
	matriz = make([][]rune, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&basura)
		matriz[i] = []rune(basura)
	}

	minimo := INF
	for k := 0; k < 4; k++ {
		actual = dijkstra(0, 0, k)
		if actual < minimo {
			fmt.Println(k, actual)
			minimo = actual
		}
	}
	fmt.Println(minimo)

}
