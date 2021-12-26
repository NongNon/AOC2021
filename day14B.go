package main7

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getMap(lines []string) map[string]string {
	pairInsersions := make(map[string]string)
	for _, line := range lines {
		pair := strings.Split(line, " -> ")
		pairInsersions[pair[0]] = pair[1]
	}
	return pairInsersions
}

func main() {
	bytes, _ := ioutil.ReadFile("in.txt")
	lines := strings.Split(string(bytes), "\n")

	polyTemplate := lines[0]
	pairInsersions := getMap(lines[2:])

	pairCount := map[string]int{}

	for i := 0; i < len(polyTemplate)-1; i++ {
		pairCount[polyTemplate[i:i+2]]++
	}

	step := 40
	for t := 0; t < step; t++ {
		newPairCount := map[string]int{}
		for key, value := range pairCount {
			newPairCount[string(key[0])+pairInsersions[key]] += value
			newPairCount[pairInsersions[key]+string(key[1])] += value
		}
		pairCount = newPairCount
	}
	//fmt.Println(pairCount)

	charCount := map[string]uint{}
	for key, value := range pairCount {
		charCount[string(key[0])] += uint(value)
	}
	charCount[string(polyTemplate[len(polyTemplate)-1])]++
	//fmt.Println(charCount)

	max, min := uint(0), uint(^uint(0))
	for _, v := range charCount {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)
}
