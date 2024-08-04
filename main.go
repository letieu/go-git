package main

import (
	"time"
)

type Hash string

type Commit struct {
	Hash       Hash
	Children   []Hash
	AuthorDate time.Time
}

// The order is determined by the relationship between commits.
// If B is child of A, B should be placed before A. (B <- A)
func TopologicalSort(commits map[Hash]Commit, newest []Hash) []Hash {
	visited := make(map[Hash]bool)
	var order []Hash

	var visit func(hash Hash)
	visit = func(hash Hash) {
		if visited[hash] {
			return
		}

		visited[hash] = true
		for _, child := range commits[hash].Children {
			visit(child)
		}

		order = append(order, hash)

	}

	for _, hash := range newest {
		visit(hash)
	}

	return order
}
