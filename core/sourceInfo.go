package core

type Category struct {
	Type string
	Urls []string
}

type Source struct {
	Id       int
	Country  string
	Vertical int //Valid[1,2,4]
	Url      string
	Category []Category
}

const (
	FOR_SALE = "FOR_SALE"
	FOR_RENT = "FOR_RENT"
)
