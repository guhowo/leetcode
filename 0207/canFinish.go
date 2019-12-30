package _207

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for i, _ := range adj {
		adj[i] = make([]int, 0)
	}

	//统计每门课程的入度以及邻接表，即被依赖的情况
	for _, w := range prerequisites {
		for _, v := range w[1:] {
			indegree[v]++
			adj[w[0]] = append(adj[w[0]], v)
		}
	}

	list := make([]int, 0)

	for i, _ := range indegree {
		if indegree[i] == 0 {
			list = append(list, i)
		}
	}

	var p int
	for len(list) != 0 {
		p = list[0]
		fmt.Println(p)
		for _, depend := range adj[p] {
			indegree[depend]--
			if indegree[depend] == 0 {
				list = append(list, depend)
			}
		}
		list = list[1:]
	}

	for _, v := range indegree {
		if v != 0 {
			return false
		}
	}
	return true
}
