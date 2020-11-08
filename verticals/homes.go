package verticals

const (
	SQUARED_METERS = "m2"
	CURRENCY_EUR   = "EUR"
)

type Price struct {
	Value    string
	Currency string
}

type FloorArea struct {
	Value string
	Unit  string
}
type PlotArea struct {
	Value string
	Unit  string
}
type Homes struct {
	Id           string
	Url          string
	Title        string
	Content      string
	Type         string
	PropertyType string
	Price        Price
	FloorArea    FloorArea
	PlotArea     PlotArea
	City         string
	Region       string
	Postcode     string
	Adress       string
	Longitude    string
	Latitude     string
	Rooms        string
	Bathrooms    string
	Pictures     [][]string
}
