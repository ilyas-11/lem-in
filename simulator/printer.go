package simulator

import (
	"fmt"
	"strings"

	"lem-in/types"
)

// Printpaths visualizes the movement of ants through the farm
// Shows step-by-step progress of each ant moving through their assigned paths
// Outputs the movements in the format "L<ant_number>-<room_name>"
func Printpaths(data *types.FarmData, Paths *types.Path) {
	paths := Paths.Pathgroup[Paths.Bestgroup]
	dants := Paths.Dant

	maxSteps := 0
	for pathIdx, path := range paths {
		if pathIdx < len(dants) {
			steps := len(path) + len(dants[pathIdx]) - 1
			if steps > maxSteps {
				maxSteps = steps
			}
		}
	}

	for step := 1; step <= maxSteps; step++ {
		var movements []string
		for pathIdx, path := range paths {
			if pathIdx >= len(dants) {
				continue
			}
			for antIdx, ant := range dants[pathIdx] {
				roomPos := step - antIdx - 1
				if roomPos >= 0 && roomPos < len(path) {
					movements = append(movements, fmt.Sprintf("%s-%s", ant, path[roomPos]))
				}
			}
		}
		if len(movements) > 0 {
			fmt.Println(strings.Join(movements, ""))
		}
	}
}
