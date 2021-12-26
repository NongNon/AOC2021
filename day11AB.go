package main1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func bloom(arr [][]int, i, j int) int {
	around := [8][2]int{
		{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1},
		{i, j - 1}, {i, j + 1},
		{i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1},
	}
	count := 1
	for _, a := range around {
		if a[0] >= 0 && a[0] < len(arr) && a[1] >= 0 && a[1] < len(arr) {
			arr[a[0]][a[1]]++
			if arr[a[0]][a[1]] == 10 {
				count += bloom(arr, a[0], a[1])
			}
		}
	}
	return count
}

func main() {
	defer writer.Flush()

	input := 10
	arr := make([][]int, input)
	for i := range arr {
		arr[i] = make([]int, input)
	}
	for i := 0; i < input; i++ {
		var n string
		scanf("%s\n", &n)
		for index, element := range strings.Split(n, "") {
			arr[i][index], _ = strconv.Atoi(element)
		}
	}
	count := 0
	day := 500
	for d := 0; d < day; d++ {
		for i := 0; i < input; i++ {
			for j := 0; j < input; j++ {
				arr[i][j]++
				if arr[i][j] == 10 {
					count += bloom(arr, i, j)
				}
			}
		}
		isEvery := true
		for i := 0; i < input; i++ {
			for j := 0; j < input; j++ {
				if arr[i][j] >= 10 {
					arr[i][j] = 0
				}
				if isEvery && arr[i][j] != 0 {
					isEvery = false
				}
			}
		}

		if isEvery {
			printf("%d\n", d+1)
			return
		}
	}

	for i := 0; i < input; i++ {
		for j := 0; j < input; j++ {
			printf("%d ", arr[i][j])
		}
		printf("\n")
	}
	printf("%d\n", count)
}
