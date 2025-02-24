package main

import "sort"

func main() {

}

func topKFrequent(words []string, k int) []string {
	unique := []string{}
	frequent := make(map[string]int)
	for _, w := range words {
		if _, ok := frequent[w]; ok {
			frequent[w]++
		} else {
			unique = append(unique, w)
			frequent[w] = 0
		}

	}

	sort.SliceStable(unique, func(i, j int) bool {
		if frequent[unique[i]] == frequent[unique[j]] {
			return unique[i] < unique[j]
		}
		return frequent[unique[i]] > frequent[unique[j]]
	})

	return unique[:k]
}
