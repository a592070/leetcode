package main

import (
	"fmt"
	"sort"
)

func main() {
	anagrams := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(anagrams)
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	} else if len(strs) == 1 {
		return [][]string{[]string{strs[0]}}
	}
	bitKeys := "00000000000000000000000000"
	fmt.Println([]rune(bitKeys))
	fmt.Println([]rune("a"))
	mapping := map[string][]string{}
	for _, str := range strs {
		r := []rune(bitKeys)
		for _, b := range str {
			fmt.Println(b)
			r[b-97] = 49
		}
		bKey := string(r)
		mapping[bKey] = append(mapping[bKey], str)
	}

	rs := [][]string{}
	for _, val := range mapping {
		rs = append(rs, val)
	}
	sort.Slice(rs, func(i, j int) bool {
		return len(rs[i]) > len(rs[j])
	})
	return rs
}
