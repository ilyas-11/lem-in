package pathfinder

import (
	"lem-in/types"
)

// BFS implements Breadth-First Search algorithm to find paths from start to end room
// Takes farm data, starting room, visited rooms map, and flag indicating if direct end access exists
// Returns a slice of room names representing the found path, or nil if no path exists
func BFS(data *types.FarmData, start string, visited map[string]bool, hasDirectAccessToEnd bool) []string {
	end := data.EndRoom

	queue := []string{start}
	Source := map[string]string{}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == end {
			if len(Allpaths(data, start, Source)) == 1 && hasDirectAccessToEnd {
				continue
			} else {
				return Allpaths(data, start, Source)
			}
		}
		for _, neighbor := range data.Links[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				Source[neighbor] = current
				queue = append(queue, neighbor)
			}
		}

	}

	return nil
}

func MarkVisited(paths [][]string, data *types.FarmData) map[string]bool {
	visited := make(map[string]bool)
	for _, path := range paths {
		for _, room := range path {
			if room != data.EndRoom {
				visited[room] = true
			}
		}
	}
	return visited
}

func Allpaths(data *types.FarmData, start string, Source map[string]string) []string {
	var path []string
	for at := data.EndRoom; at != ""; at = Source[at] {

		if at == data.StartRoom {
			break
		}
		path = append([]string{at}, path...)
		if at == start {
			break
		}
	}
	return path
}
