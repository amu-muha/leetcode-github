func countServers(grid [][]int) int {

    m := len(grid)
    n := len(grid[0])

    rowCount := make([]int, m)
    colCount := make([]int, n)

    // First pass: count servers in each row and column
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {

            if grid[i][j] == 1 {
                rowCount[i]++
                colCount[j]++
            }
        }
    }

    // Second pass: count communicating servers
    answer := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {

            if grid[i][j] == 1 {
                if rowCount[i] > 1 || colCount[j] > 1 {
                    answer++
                }
            }
        }
    }

    return answer
}