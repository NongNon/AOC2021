package main4

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func createGraph(points []string) [][]bool {
	max_x, max_y := 0, 0
	for _, point := range points {
		temp := strings.Split(point, ",")
		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])
		if x > max_x {
			max_x = x
		}
		if y > max_y {
			max_y = y
		}
	}

	graph := make([][]bool, max_y+1)
	for i := range graph {
		graph[i] = make([]bool, max_x+1)
	}

	for _, point := range points {
		temp := strings.Split(point, ",")
		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])
		graph[y][x] = true
	}

	return graph
}

func fold(graph [][]bool, axis rune, n int) [][]bool {
	if axis == 'x' {
		for row := 0; row < len(graph); row++ {
			for foldDist := 1; foldDist <= n; foldDist++ {
				graph[row][n-foldDist] = graph[row][n-foldDist] || graph[row][n+foldDist]
			}
			graph[row] = graph[row][:n]
		}
	} else {
		for col := 0; col < len(graph[0]); col++ {
			for foldDist := 1; foldDist <= n; foldDist++ {
				graph[n-foldDist][col] = graph[n-foldDist][col] || graph[n+foldDist][col]
				graph[n+foldDist][col] = false
			}
		}
		return graph[:n]
	}
	return graph
}

func main() {
	bytes, _ := ioutil.ReadFile("in.txt")
	input := string(bytes)
	lines := strings.Split(input, "\n")

	n_points := 853
	graph := createGraph(lines[:n_points])

	graph = fold(graph, 'x', 655)

	count := 0
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
	//print graph
	// for i := range graph {
	// 	for j := range graph[i] {
	// 		if graph[i][j] {
	// 			fmt.Printf("#")
	// 		} else {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }
}
