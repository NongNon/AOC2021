package main6

import (
	"io/ioutil"
	"strings"
)

func getMap(lines []string) map[string]byte {
	pairInsersions := make(map[string]byte)
	for _, line := range lines {
		pair := strings.Split(line, " -> ")
		pairInsersions[pair[0]] = pair[1][0]
	}
	return pairInsersions
}

func dfs(a, b byte, pairInsersion map[string]byte, charCount map[byte]int, step int) {
	if step == 0 {
		return
	}
	insert := pairInsersion[string(a)+string(b)]
	charCount[insert]++
	dfs(a, insert, pairInsersion, charCount, step-1)
	dfs(insert, b, pairInsersion, charCount, step-1)
}

func main() {
	bytes, _ := ioutil.ReadFile("in.txt")
	lines := strings.Split(string(bytes), "\n")
	charCount := map[byte]int{}

	polyTemplate := lines[0]
	pairInsersions := getMap(lines[2:])

	step := 10
	for i := 0; i < len(polyTemplate); i++ {
		charCount[polyTemplate[i]]++
	}
	for i := 0; i < len(polyTemplate)-1; i++ {
		dfs(polyTemplate[i], polyTemplate[i+1], pairInsersions, charCount, step)
	}

	max, min := 0, 99999999
	for _, v := range charCount {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	println(max - min)
	//print charCount
	// for k, v := range charCount {
	// 	fmt.Printf("%c: %d\n", k, v)
	// }
}
