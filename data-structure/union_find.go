package main

//when to use:
//	- finding number of component groups / connected components
// example: https://leetcode.com/problems/accounts-merge

type UnionFind struct {
	Root map[string]string
	Rank map[string]int
}

func (u *UnionFind) Find(s string) string {
	// // quick union, time O(n)
	// for s != u.Root[s] {
	//     s = u.Root[s]
	// }
	// return s

	// find with path compression, time O(log n)
	if s == u.Root[s] {
		return s
	}
	u.Root[s] = u.Find(u.Root[s])
	return u.Root[s]
}

func (u *UnionFind) Union(s1, s2 string) {
	// // quick union, time O(n)
	// s1 = u.Find(s1)
	// s2 = u.Find(s2)
	// u.Root[s1] = s2

	// union by rank, time O(log n), this keeps the height as low as possible
	s1 = u.Find(s1)
	s2 = u.Find(s2)
	if s1 == s2 {
		return
	}

	// assign the lower rank to higher rank
	if u.Rank[s1] > u.Rank[s2] {
		u.Root[s2] = s1
	} else if u.Rank[s1] < u.Rank[s2] {
		u.Root[s1] = s2
	} else {
		u.Root[s2] = s1
		u.Rank[s1]++
	}
}

func (u *UnionFind) FindUniqueRoots() int {
	// time O(n log n) -> access all n elements and log n for finding the root for each element
	var uniqueRoots int
	found := make(map[string]bool)
	for _, v := range u.Root {
		v = u.Find(v)
		if !found[v] {
			found[v] = true
			uniqueRoots++
		}
	}

	return uniqueRoots
}

func NewUF(strs []string) *UnionFind {
	root := make(map[string]string)

	for _, s := range strs {
		root[s] = s
	}

	return &UnionFind{Root: root, Rank: make(map[string]int)}
}
