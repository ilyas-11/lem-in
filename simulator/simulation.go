package simulator

import (
	"lem-in/types"
	"strconv"
)

// Simulation simulates ant movement through the farm's paths
// It assigns ants to different path groups and finds the most efficient distribution
// that minimizes the total number of moves needed to get all ants to the end
func Simulation(data *types.FarmData, Paths *types.Path) {
	allgroups := Paths.Pathgroup
	longegroup := 0

	for i, pathgroup := range allgroups {
		ant := data.AntsCount
		dant := make([][]string, len(pathgroup))
		for ant > 0 {
			for j := 0; j < len(pathgroup); j++ {
				if j == len(pathgroup)-1 && len(pathgroup[j])+len(dant[j]) > len(pathgroup[0])+len(dant[0]) {
					continue
				}
				if j+1 < len(pathgroup) && j+1 < len(dant) && len(pathgroup[j])+len(dant[j]) > len(pathgroup[j+1])+len(dant[j+1]) {
					continue
				}
				if ant <= 0 {
					break
				}

				dant[j] = append(dant[j], "L"+strconv.Itoa(data.AntsCount-ant+1))
				if len(dant[j]) > 0 {
					Paths.Pathused = append(Paths.Pathused, j)
				}
				ant--
			}
		}

		longe := len(pathgroup[0])
		if i == 0 {
			longegroup = longe + len(dant[0])
			Paths.Dant = dant
		}

		for j, path := range pathgroup {
			if len(path)+len(dant[j]) >= longe {
				longe = len(path) + len(dant[j])
			}

		}

		if longegroup > longe {
			Paths.Bestgroup = i
			longegroup = longe
			Paths.Dant = dant
		}

	}

}
