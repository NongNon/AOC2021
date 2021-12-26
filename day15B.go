package main9

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Item struct {
	dist int
	pair [2]int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func main() {
	bytes, _ := ioutil.ReadFile("in.txt")
	lines := strings.Split(string(bytes), "\n")

	arr := make([][]uint16, len(lines)*5)
	visited := make([][]bool, len(lines)*5)
	for i, line := range lines {
		for row := 0; row < 5; row++ {
			x := i + row*len(lines)
			arr[x] = make([]uint16, len(line)*5)
			visited[x] = make([]bool, len(line)*5)
		}
		for j, c := range line {
			for row := 0; row < 5; row++ {
				x := i + row*len(lines)
				for col := 0; col < 5; col++ {
					y := j + col*len(line)

					weight := uint16((int(c) - 48 + row + col) % 9)
					if weight == 0 {
						weight = 9
					}
					arr[x][y] = weight //+ uint16(len(lines)*5-x-1) + uint16(len(lines)*5-y-1)
				}
			}
		}
	}
	arr[0][0] = 0

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Item{dist: 0, pair: [2]int{0, 0}})

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Item)
		dist, i, j := node.dist, node.pair[0], node.pair[1]
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		if i == len(arr)-1 && j == len(arr[i])-1 {
			fmt.Println(dist)
			break
		}

		around := [4][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
		for _, p := range around {
			if p[0] < 0 || p[1] < 0 || p[0] >= len(arr) || p[1] >= len(arr[p[0]]) {
				continue
			}
			if visited[p[0]][p[1]] {
				continue
			}
			heap.Push(&pq, &Item{dist: dist + int(arr[p[0]][p[1]]), pair: p})
		}
	}

	//print arr
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }
}
