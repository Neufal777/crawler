package verticals

type Price struct {
	Value    int
	Currency string
}

type FloorArea struct {
	Value int
	Unit  string
}
type PlotArea struct {
	Value int
	Unit  string
}
type Homes struct {
	Id           int
	Title        string
	Content      string
	Type         string
	PropertyType string
	price        Price
	FloorArea    FloorArea
	PlotArea     PlotArea
}
