package core

import (
	"log"
	"strconv"
)

/*
Categories is used for getting all the categories from
the Url to parse all the ads inside, including the pagination
*/

func ClassifyCat(cat *Source) ([]string, []string) {

	Sale := []string{}
	Rent := []string{}

	/*
		Separate all the category urls between rent and sale

		Sale []string - will contain all FOR_SALE categories
		Rent []string - will contain all FOR_RENT categories
	*/
	for _, cat := range cat.Category {

		switch cat.Type {
		case "FOR_SALE":
			for _, url := range cat.Urls {
				Sale = append(Sale, url)
			}
		case "FOR_RENT":
			for _, url := range cat.Urls {
				Rent = append(Rent, url)
			}
		}
	}

	/*
		Once we have 2 slices with all the categories of each,
		we send them to be processed
	*/

	return Sale, Rent

}

func Pagination(catUrls []string) {

	/*
		With a specific pattern, all the categories pages are saved
		to be processed afterwards
	*/

	for _, catUrl := range catUrls {

		/*
			Foreach categorie url, make the pagination
		*/
		for i := 0; i < 6; i++ {
			log.Println(catUrl + "&page=" + strconv.Itoa(i))
		}
	}
}
