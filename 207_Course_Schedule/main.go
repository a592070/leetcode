package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	prereq := map[int][]int{}
	for _, pre := range prerequisites {
		// skip obvious cycle
		if pre[0] == pre[1] {
			return false
		}
		// put node's parents together
		prereq[pre[0]] = append(prereq[pre[0]], pre[1])
	}

	// check whether struggle in a cycle
	visiting := map[int]bool{}
	var dfs func(course int) bool
	dfs = func(course int) bool {
		// struggled in cycle
		if visiting[course] {
			return false
		}

		if len(prereq[course]) == 0 {
			return true
		}

		visiting[course] = true

		// check nod's parent is reachable
		for _, pre := range prereq[course] {
			if !dfs(pre) {
				return false
			}
		}

		visiting[course] = false

		// skip visited node
		prereq[course] = []int{}
		return true
	}

	// check all nodes are reachable
	for i := 0; i <= numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true

}
