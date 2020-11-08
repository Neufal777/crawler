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
	ID       = 281
	COUNTRY  = "ES"
	VERTICAL = 1
	URL      = "https://www.pisos.com/"
	AD_REGEX = "data-navigate-ref=.([^\"']*)"
)
