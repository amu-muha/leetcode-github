import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
    sort.Slice(logs, func(i, j int) bool {
        return logs[i][1] < logs[j][1]
    })

    type Query struct {
        time  int
        index int
    }

    sortedQueries := make([]Query, len(queries))

    for i, q := range queries {
        sortedQueries[i] = Query{q, i}
    }

    sort.Slice(sortedQueries, func(i, j int) bool {
        return sortedQueries[i].time < sortedQueries[j].time
    })

    answer := make([]int, len(queries))
    freq := make(map[int]int)

    left, right := 0, 0
    active := 0

    for _, q := range sortedQueries {
        start := q.time - x
        end := q.time

        // Add logs that fall before or at the query's end time.
        for right < len(logs) && logs[right][1] <= end {
            serverID := logs[right][0]

            if freq[serverID] == 0 {
                active++
            }

            freq[serverID]++
            right++
        }

        // Remove logs that fall before the query's start time.
        for left < right && logs[left][1] < start {
            serverID := logs[left][0]

            freq[serverID]--

            if freq[serverID] == 0 {
                active--
            }

            left++
        }

        answer[q.index] = n - active
    }

    return answer
}