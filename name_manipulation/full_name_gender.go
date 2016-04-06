package main

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"os"
	"strings"
)

type Node struct {
	Children map[rune]*Node
	Gender   *string
}

type Trie struct {
	root Node
}

func (t *Trie) insert(node *Node, word []rune, gender string) {
	if len(word) == 0 {
		node.Gender = &gender
	} else {
		_, ok := node.Children[word[0]]
		if !ok {
			node.Children[word[0]] = &Node{Children: make(map[rune]*Node)}
		}
		t.insert(node.Children[word[0]], word[1:], gender)
	}
}

func (t *Trie) query(node *Node, word []rune) *string {

	if len(word) == 0 {
		return node.Gender
	}

	_, ok := node.Children[word[0]]
	if !ok {
		return node.Gender
	}
	x := t.query(node.Children[word[0]], word[1:])
	if x == nil {
		return node.Gender
	}
	return x
}

func (t *Trie) Query(word []rune) *string {
	return t.query(&t.root, word)
}

func (t *Trie) Insert(word []rune, gender string) {
	t.insert(&t.root, word, gender)
}

func main() {

	trie := Trie{root: Node{Children: make(map[rune]*Node)}}
	data, err := ioutil.ReadFile("./used_genders.txt")
	if err != nil {
		panic(err)
	}
	c, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	if err != nil {
		panic(err)
	}

	for i, v := range c {
		if i == 0 || len(v) < 3 {
			continue
		}
		v[0] = strings.ToLower(v[0])
		trie.Insert([]rune(v[0]), v[1])
	}

	// Find names
	data, _ = ioutil.ReadFile("./names.txt")
	c, _ = csv.NewReader(bytes.NewReader(data)).ReadAll()

	out, _ := os.Create("./output.txt")
	csvout := csv.NewWriter(out)

	csvout.Write([]string{"name", "gender"})
	for i, v := range c {
		if i == 0 {
			continue
		}
		tmp := strings.ToLower(v[1])
		g := trie.Query([]rune(tmp))
		var gender string
		if g == nil {
			gender = ""
		} else {
			gender = *g
		}
		csvout.Write([]string{v[1], gender})
	}
	csvout.Flush()
}
