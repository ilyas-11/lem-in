package pathfinder

import (
	"fmt"
	"lem-in/types"
)

// Grouppaths analyzes the ant farm data and finds groups of possible paths for ants to traverse
// The function performs the following:
// 1. Creates initial paths from the start room through each available link
// 2. For each initial path, checks if there's direct access to the end room
// 3. Looks for additional sub-paths that can be used in parallel while avoiding path conflicts

// 4. Groups compatible paths together to allow multiple ants to move simultaneously
//
// Parameters:
//   - data: Pointer to FarmData containing rooms, links and ant farm configuration
//
// Returns:
//   - Pointer to Path structure containing groups of valid paths from start to end
func Grouppaths(data *types.FarmData) (*types.Path, error) {
	Paths := &types.Path{
		Pathgroup: make([][][]string, 0),
	}
	isempty := false
	for _, links := range data.Links[data.StartRoom] {
		gpath := [][]string{}
		gpath = append(gpath, BFS(data, links, map[string]bool{data.StartRoom: true}, false))
		if len(gpath[0]) == 0 {
			continue
		}
		Paths.Pathgroup = append(Paths.Pathgroup, gpath)
		if len(gpath[0]) != 0 {
			isempty = true
		}
	}
	if !isempty {
		return nil, fmt.Errorf("Error: no paths found from start to end")
	}

	for i, path := range Paths.Pathgroup {

		var hasDirectAccessToEnd bool

		if len(path[0]) == 1 {
			hasDirectAccessToEnd = true
		}

		visited := MarkVisited(path, data)

		for _, links := range data.Links[data.StartRoom] {
			if !visited[links] {

				subpath := BFS(data, data.StartRoom, visited, hasDirectAccessToEnd)

				if subpath != nil {

					Paths.Pathgroup[i] = append(Paths.Pathgroup[i], subpath)
					visited = MarkVisited(Paths.Pathgroup[i], data)
				}
			}
		}
		if len(Paths.Pathgroup[i]) == 1 {
		}
	}
	return Paths, nil
}
