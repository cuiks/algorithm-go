package main

import "sort"

func findItinerary(tickets [][]string) []string {
	result := make([]string, 0)
	record := map[string][]string{}
	for _, v := range tickets {
		record[v[0]] = append(record[v[0]], v[1])
	}

	for k := range record {
		sort.Strings(record[k])
	}

	var dfs func(cur string)
	dfs = func(cur string) {
		for {
			if v, ok := record[cur]; !ok || len(v) == 0 {
				break
			}
			tmp := record[cur][0]
			record[cur] = record[cur][1:]
			dfs(tmp)
		}
		result = append(result, cur)
	}
	dfs("JFK")

	i, j := 0, len(result)-1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}

func main() {

}
