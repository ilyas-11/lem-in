package parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem-in/types"
)

// ParseFile reads and parses an ant farm configuration file
// filename: path to the configuration file
// returns: pointer to FarmData structure containing the parsed data and any error encountered
func ParseFile(filename string) (*types.FarmData, error) {
	data := &types.FarmData{
		AntsCount: -1,
		Rooms:     make([]string, 0),
		Links:     make(map[string][]string),
		Roomloc:   make(map[string][2]int),
	}
	state := &types.ParserState{
		HasLinks:       false,
		HasStartLinks:  false,
		HasEndLinks:    false,
		InRoomsSection: false,
		ExpectStart:    false,
		ExpectEnd:      false,
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			if line == "" {
				continue
			}
			if state.ExpectStart || state.ExpectEnd {
				return nil, fmt.Errorf("Line %d: Expected room definition after ##start/##end directive but found none", i+1)
			}
			if line == "##start" {
				state.ExpectStart = true
			} else if line == "##end" {
				state.ExpectEnd = true
			}
			continue
		}

		if data.AntsCount == -1 {
			ants, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("Line %d: Invalid ant count - must be a valid number", i+1)
			}
			if ants <= 0 {
				return nil, fmt.Errorf("Invalid ant count: must be greater than 0")
			}
			data.AntsCount = ants
			state.InRoomsSection = true
		} else if strings.Contains(line, "-") {
			state.InRoomsSection = false
			if !state.InRoomsSection { //+
				parts := strings.Split(line, "-")
				if len(parts) != 2 {
					return nil, fmt.Errorf("Line %d: Invalid link format - expected 'room1-room2'", i+1)
				}
				if _, ok := data.Roomloc[parts[0]]; !ok {
					return nil, fmt.Errorf("Line %d: Room '%s' not found - must define rooms before linking them", i+1, parts[0])
				}
				if _, ok := data.Roomloc[parts[1]]; !ok {
					return nil, fmt.Errorf("Line %d: Room '%s' not found - must define rooms before linking them", i+1, parts[1])
				}
				if _, ok := data.Links[parts[0]]; ok {
					for _, linkedRoom := range data.Links[parts[0]] {
						if linkedRoom == parts[1] {
							return nil, fmt.Errorf("Line %d: Link between '%s' and '%s' already exists", i+1, parts[0], parts[1])
						}
					}
				}
				if _, ok := data.Links[parts[1]]; ok {
					for _, linkedRoom := range data.Links[parts[1]] {
						if linkedRoom == parts[0] {
							return nil, fmt.Errorf("Line %d: Link between '%s' and '%s' already exists", i+1, parts[1], parts[0])
						}
					}
				}
				if parts[0] == parts[1] {
					return nil, fmt.Errorf("Line %d: Cannot link room '%s' to itself", i+1, parts[0])
				}
				if  parts[0] == data.StartRoom || parts[1] == data.StartRoom {
					state.HasStartLinks = true
				}
				if  parts[0] == data.EndRoom || parts[1] == data.EndRoom {
					state.HasEndLinks = true
				}

				state.HasLinks = true
				data.Links[parts[0]] = append(data.Links[parts[0]], parts[1])
				data.Links[parts[1]] = append(data.Links[parts[1]], parts[0])
			} //+
		} else {
			parts := strings.Fields(line)

			if len(parts) == 3 {
				if !state.InRoomsSection {
					return nil, fmt.Errorf("line %d: invalid file structure - all room definitions must come before links", i+1)
				}
				if parts[0][0] == 'L' {
					return nil, fmt.Errorf("line %d: room name cannot start with 'L' (reserved for ant names)", i+1)
				}
				x, errX := strconv.Atoi(parts[1])
				y, errY := strconv.Atoi(parts[2])
				if errX != nil || errY != nil {
					return nil, fmt.Errorf("line %d: invalid room coordinates - must be valid numbers", i+1)
				}
				if x < 0 || y < 0 {
					return nil, fmt.Errorf("line %d: room coordinates must be non-negative values", i+1)
				}
				for _, room := range data.Rooms {
					if room == parts[0] {
						return nil, fmt.Errorf("line %d: room name '%s' already exists - room names must be unique", i+1, parts[0])
					}
				}
				for _, coord := range data.Roomloc {
					if coord == [2]int{x, y} {
						return nil, fmt.Errorf("line %d: coordinates (%d, %d) already used by another room - room positions must be unique", i+1, x, y)
					}
				}
				if state.ExpectStart {
					data.StartRoom = parts[0]
					state.ExpectStart = false

				} else if state.ExpectEnd {
					data.EndRoom = parts[0]
					state.ExpectEnd = false

				}
				data.Roomloc[parts[0]] = [2]int{x, y}
				data.Rooms = append(data.Rooms, parts[0])
			} else {
				return nil, fmt.Errorf("line %d: invalid room format - expected 'name x y' where x and y are coordinates", i+1)
			}

		}
	}
	if data.StartRoom == "" || data.EndRoom == "" {
		return nil, fmt.Errorf("missing required ##start and/or ##end room definitions")
	} else if !state.HasLinks {
		return nil, fmt.Errorf("no links found - farm must have at least one path between rooms")
	}else if !state.HasStartLinks {
		return nil, fmt.Errorf("no links found for start room '%s' - start room must have at least one connection", data.StartRoom)
	}else if !state.HasEndLinks {
		return nil, fmt.Errorf("no links found for end room '%s' - end room must have at least one connection", data.EndRoom)
	}
	fmt.Println(string(content))
	return data, nil
}
