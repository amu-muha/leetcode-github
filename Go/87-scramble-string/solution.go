func isScramble(s1 string, s2 string) bool {
	memo := make(map[string]bool)

	var dfs func(a, b string) bool
	dfs = func(a, b string) bool {
		if a == b {
			return true
		}

		key := a + "#" + b
		if val, exists := memo[key]; exists {
			return val
		}

		n := len(a)

	
		var count [26]int
		for i := 0; i < n; i++ {
			count[a[i]-'a']++
			count[b[i]-'a']--
		}
		for _, c := range count {
			if c != 0 {
				memo[key] = false
				return false
			}
		}

		
		for i := 1; i < n; i++ {
			
			if dfs(a[:i], b[:i]) && dfs(a[i:], b[i:]) {
				memo[key] = true
				return true
			}

			
			if dfs(a[:i], b[n-i:]) && dfs(a[i:], b[:n-i]) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false
		return false
	}

	return dfs(s1, s2)
}