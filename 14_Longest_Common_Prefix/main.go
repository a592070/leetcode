package main

/*
https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/
func main() {

}

type Trie struct {
	Root *Node
}
type Node struct {
	IsEnd    bool
	Value    string
	Children map[string]*Node
}

func NewPrefixTrie() *Trie {
	return &Trie{
		Root: &Node{
			Value:    "",
			Children: make(map[string]*Node),
		},
	}
}

func (t *Trie) Insert(word string) {
	temp := t.Root
	for _, char := range word {
		charStr := string(char)
		if _, ok := temp.Children[charStr]; !ok {
			temp.Children[charStr] = &Node{
				Value:    charStr,
				Children: make(map[string]*Node),
			}
		}
		temp = temp.Children[charStr]
	}
}

func getMiniStr(strs []string) string {
	ministr := strs[0]
	for _, char := range strs {
		if len(string(char)) < len(ministr) {
			ministr = string(char)
		}
	}
	return ministr
}

func longestCommonPrefix_Tier(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var result string

	prefixTrie := NewPrefixTrie()

	for _, el := range strs {
		prefixTrie.Insert(el)
	}

	temp := prefixTrie.Root
	minStr := getMiniStr(strs)

	for _, el := range minStr {
		strEl := string(el)

		if len(temp.Children) == 1 {
			result += strEl
			temp = temp.Children[strEl]
		} else {
			return result
		}
	}

	return result
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || str[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
}
