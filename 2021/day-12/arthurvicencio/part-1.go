package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	Name     string
	Children []*Node
}

func main() {

	inputFile, _ := ioutil.ReadFile("input.txt")

	parsedInput := strings.Split(string(inputFile), "\n")

	nodes := make(map[string]*Node)
	for _, cardInput := range parsedInput {
		line := strings.Split(cardInput, "-")

		from := line[0]
		to := line[1]
		if _, ok := nodes[from]; !ok {
			nodes[from] = &Node{Name: from, Children: make([]*Node, 0)}
		}
		if _, ok := nodes[to]; !ok {
			nodes[to] = &Node{Name: to, Children: make([]*Node, 0)}
		}
		nodes[from].Children = append(nodes[from].Children, nodes[to])
		nodes[to].Children = append(nodes[to].Children, nodes[from])
	}

	type Path struct {
		Node *Node
		Seen map[string]int
	}

	var count int

	queue := list.New()
	queue.PushBack(Path{nodes["start"], make(map[string]int)})

	for queue.Len() > 0 {

		elem := queue.Front()
		queue.Remove(elem)

		path := elem.Value.(Path)

		if path.Seen[path.Node.Name] >= 1 && isLower(path.Node.Name) {
			continue
		}
		path.Seen[path.Node.Name]++

		if path.Node.Name == "end" {
			count++
		}

		for _, child := range path.Node.Children {
			queue.PushBack(Path{Node: child, Seen: copyMap(path.Seen)})
		}
	}

	fmt.Println(count)
}

func isLower(s string) bool {
	return s == strings.ToLower(s)
}

func copyMap(m map[string]int) map[string]int {
	n := make(map[string]int)
	for i, v := range m {
		n[i] = v
	}
	return n
}
