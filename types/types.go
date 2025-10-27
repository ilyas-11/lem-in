package types

type FarmData struct {
	AntsCount int
	Rooms     []string
	Roomloc   map[string][2]int
	Links     map[string][]string
	StartRoom string
	EndRoom   string
}
type Path struct {
	Pathgroup [][][]string
	Bestgroup int
	Dant      [][]string
	Pathused  []int
}
type ParserState struct {
	HasLinks, HasStartLinks, HasEndLinks bool
	InRoomsSection                       bool
	ExpectStart, ExpectEnd               bool
}
