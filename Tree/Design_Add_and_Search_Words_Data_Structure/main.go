package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AddWord("abc")
	param_2 := obj.Search("a.c")
	fmt.Println(param_2)
}

type Node struct {
	char     byte
	hasWord  bool
	children []*Node
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{
			children: make([]*Node, 26),
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	current := this.root
	for i := 0; i < len(word); i++ {
		val := current.children[word[i]-'a']
		if val == nil {
			current.children[word[i]-'a'] = &Node{
				char:     word[i],
				hasWord:  false,
				children: make([]*Node, 26),
			}
		}

		current = current.children[word[i]-'a']
	}

	current.hasWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.SearchRec(this.root, word, 0)
}

func (this *WordDictionary) SearchRec(current *Node, word string, index int) bool {
	if current == nil {
		return false
	}

	if index == len(word) {
		return current.hasWord
	}

	if word[index] == '.' {
		for _, child := range current.children {
			r := this.SearchRec(child, word, index+1)
			if r {
				return true
			}
		}
		return false
	} else {
		val := current.children[word[index]-'a']
		if val == nil {
			return false
		}

		return this.SearchRec(val, word, index+1)
	}
}
