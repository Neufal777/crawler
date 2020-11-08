package core

import (
	"log"
	"regexp"
)

/*
Categories is used for getting all the categories from
the Url to parse all the ads inside, including the pagination
*/

func (cat *Source) ClassifyCat() ([]string, []string) {

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
		we send them to be processed [PAGINATION]
	*/

	return Sale, Rent

}

func GetAdUrls(cat []string) [][]string {

	var adUrls [][]string
	//count := 0
	for _, c := range cat {

		/*
			Here we will download each category html
			and find for all the available ads
		*/

		catHtml := DownloadHtml(c)
		log.Println("Page downloaded. processing..")

		rAd := regexp.MustCompile("data-navigate-ref=.([^\"']*)")

		adUrls = rAd.FindAllStringSubmatch(catHtml, 2)

		if adUrls == nil {
			/*
				If there's no more ads, go to the next page
			*/
			log.Println("This page has no ads or the Regex is wrong...")

		}

	}

	return adUrls
}

// func Pagination(catUrls []string) {

//
// 		With a specific pattern, all the categories pages are saved
// 		to be processed afterwards
//

// 	var pagination []string
// 	for _, catUrl := range catUrls {

//
// 			Foreach categorie url, make the pagination
//
// 		for i := 1; i <= 4; i++ {

// 			nextPage := catUrl + "/" + strconv.Itoa(i)
// 			pagination = append(pagination, nextPage)
// 			ProcessPage(nextPage)

// 		}
// 	}

// }

func CheckPageExists(page string) bool {

	/*
		This function checks if a specific page exists or not
	*/

	return false
}

func ProcessPage(page string) {

	/*
		This function will proceess each url of a page
		coming grom Pagination function and will get all
		the ads from there
	*/

	log.Println("Processing...")
	log.Println(page)
}
