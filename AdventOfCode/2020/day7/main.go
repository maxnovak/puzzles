package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BagRules struct {
	parentBag    string
	childrenBags map[string]int
}

type TreeNodes struct {
	BagType  string       `json:"bagType"`
	Count    int          `json:"count"`
	Children []*TreeNodes `json:"children"`
}

func main() {
	fmt.Println("this is day 7")
	trees := []TreeNodes{}
	bagRules := readFile()
	for _, bagRule := range bagRules {
		trees = insertToTree(bagRule.parentBag, bagRule.childrenBags, trees)
	}

	count := 0
	bagCount := 0
	for _, tree := range trees {
		if tree.BagType == "shiny gold bag" {
			continue
		}
		if canHoldGoldBag(&tree) {
			count++
		}
	}
	for _, tree := range trees {
		if tree.BagType == "shiny gold bag" {
			tree.Count = 1
			bagCount = traverseChildren(tree)
		}
	}
	//Print out formated trees
	// formattedTrees, _ := json.MarshalIndent(trees, "", " ")
	// fmt.Println(string(formattedTrees))

	fmt.Println("Part 1 solution", count)
	//remove the shiny gold bag from the count
	fmt.Println("Part 2 solution", bagCount-1)
}

func traverseChildren(tree TreeNodes) int {
	numberOfBags := 0
	for _, child := range tree.Children {
		if child.BagType == " other bag" {
			continue
		} else {
			numberOfBags += tree.Count * traverseChildren(*child)
		}
	}
	numberOfBags += tree.Count
	return numberOfBags
}

func canHoldGoldBag(tree *TreeNodes) bool {
	if tree.BagType == "shiny gold bag" {
		return true
	}

	if len(tree.Children) > 0 {
		for _, child := range tree.Children {
			result := canHoldGoldBag(child)
			if result {
				return result
			}
		}
	}
	return false
}

func insertToTree(bagColor string, allowedBags map[string]int, trees []TreeNodes) []TreeNodes {
	treeNode := TreeNodes{
		BagType: bagColor,
	}
	for key, value := range allowedBags {
		treeNode.Children = append(treeNode.Children, &TreeNodes{
			BagType: key,
			Count:   value,
		})
	}
	for _, tree := range treeNode.Children {
		for _, item := range trees {
			if item.BagType == tree.BagType {
				tree.Children = item.Children
			}
		}
	}
	for _, tree := range trees {
		appendToExistingNodes(tree, treeNode)
	}
	trees = append(trees, treeNode)
	return trees
}

func appendToExistingNodes(tree TreeNodes, newBag TreeNodes) {
	if len(tree.Children) == 0 {
		return
	}
	for _, child := range tree.Children {
		if child.BagType == newBag.BagType {
			child.Children = newBag.Children
		} else {
			appendToExistingNodes(*child, newBag)
		}
	}
}

func readFile() []BagRules {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rules []BagRules
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Replace(text, ".", "", -1)
		text = strings.Replace(text, "bags", "bag", -1)
		line := strings.Split(text, " contain ")
		textChildren := strings.Split(line[1], ", ")
		children := map[string]int{}
		for _, child := range textChildren {
			children[child[2:]], _ = strconv.Atoi(child[0:1])
		}

		rules = append(rules, BagRules{
			parentBag:    line[0],
			childrenBags: children,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rules
}
