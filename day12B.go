package main3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func dfs(graph map[string][]string, des string, path []string, isTwice bool) int {
	if des == "end" {
		fmt.Println(path)
		return 1
	}
	countPath := 0
	for _, v := range graph[des] {
		if v == "start" {
			continue
		}
		if IsUpper(v) {
			countPath += dfs(graph, v, path, isTwice)
		} else { // if IsLower(v)
			if isTwice {
				if !contains(path, v) {
					countPath += dfs(graph, v, append(path, v), true)
				}
			} else {
				if !contains(path, v) {
					countPath += dfs(graph, v, append(path, v), false)
				} else {
					countPath += dfs(graph, v, append(path, v), true)
				}
			}
		}
	}
	return countPath
}

func main() {
	defer writer.Flush()

	input := 22
	graph := make(map[string][]string)
	for i := 0; i < input; i++ {
		var str string
		scanf("%s\n", &str)
		tmp := strings.Split(str, "-")

		a, b := tmp[0], tmp[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	path := dfs(graph, "start", []string{"start"}, false)

	printf("%d\n", path)

	// for k, v := range graph {
	// 	fmt.Println(k, v)
	// }
}
