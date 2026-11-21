func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	adj := make([][]int, n)
	for i := 1; i < n; i++ {
		adj[parent[i]] = append(adj[parent[i]], i)
	}

	
	newParent := make([]int, n)
	copy(newParent, parent)

	
	lastSeen := make([]int, 26)
	for i := range lastSeen {
		lastSeen[i] = -1
	}

	
	var dfsFindParents func(u int)
	dfsFindParents = func(u int) {
		charIdx := int(s[u] - 'a')
		prevAncestor := lastSeen[charIdx]

		if prevAncestor != -1 {
			newParent[u] = prevAncestor
		}

		lastSeen[charIdx] = u
		for _, v := range adj[u] {
			dfsFindParents(v)
		}
		
		lastSeen[charIdx] = prevAncestor
	}

	dfsFindParents(0)


	finalAdj := make([][]int, n)
	for i := 1; i < n; i++ {
		finalAdj[newParent[i]] = append(finalAdj[newParent[i]], i)
	}


	answer := make([]int, n)
	var dfsCalcSizes func(u int) int
	dfsCalcSizes = func(u int) int {
		size := 1
		for _, v := range finalAdj[u] {
			size += dfsCalcSizes(v)
		}
		answer[u] = size
		return size
	}

	dfsCalcSizes(0)
	return answer
}