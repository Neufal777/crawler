package core

import (
	"log"
	"regexp"

	"github.com/draco/utils"
	"github.com/draco/verticals"
)

/*
Categories is used for getting all the categories from
the Url to parse all the ads inside, including the pagination
*/

func (cat *Source) ClassifyHomes() ([]string, []string) {

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

func GetAdUrls(cat string) ([][]string, int) {

	var adUrls [][]string
	count := 0

	/*
		Here we will download each category html
		and find for all the available ads
	*/

	//download the page content
	catHtml, _ := DownloadHtml(cat)

	log.Println("Page downloaded. processing..")

	//get all the ads from that page
	rAd := regexp.MustCompile(AD_REGEX) //REGEX can be found in the sourceInfo.go

	//save all the results
	adUrls = rAd.FindAllStringSubmatch(catHtml, -1)

	count += len(adUrls)

	if adUrls == nil {
		/*
			If there's no more ads, go to the next page
		*/
		log.Println("This page has no ads or the Regex is wrong...")

	}

	log.Println("Processed", count, "ads!")
	return adUrls, count
}

func ProcessAdList(adList [][]string, count int) {

	var finalAdUrls []string

	/*
		This function will proceess each ad url to:
		download it and get all the content from there
	*/

	log.Println("number of ads:", count)
	for i := 1; i < count; i++ {

		adurl := utils.CompleteUrl("https://www.pisos.com", adList[i][1])
		finalAdUrls = append(finalAdUrls, adurl)
	}

	//once we have all the ads URL we process them to get all the data
	AdExtractor(finalAdUrls)

	//log.Println(finalAdUrls)

}

func AdExtractor(ads []string) {

	/*
		Here we download the ad HTML and process them
	*/

	for _, ad := range ads {

		/*
			We download each ad html and process it to get all the fields we need
		*/

		content, err := DownloadHtml(ad)

		if err != nil {
			log.Panic(err)
		}

		//declare all the regex we are going to need to extract all that information
		var (
			rTitle     = regexp.MustCompile("<meta property=.og:title. content=.([^\"']*)")
			rContent   = regexp.MustCompile("<meta property=.og:description. content=.([^\"']*)")
			rprice     = regexp.MustCompile("h1 jsPrecioH1.>([^<]*)")
			rFloorArea = regexp.MustCompile("iv class=.icon icon-superficie.></div>([^ ]*)")
			rPlotArea  = regexp.MustCompile("iv class=.icon icon-superficie.></div>([^ ]*)")
			rRooms     = regexp.MustCompile("iv class=.icon icon-habitaciones.></div>([^<]*)")
			rBathrooms = regexp.MustCompile("iv class=.icon icon-banyos.></div>([^<]*)")
			rPictures  = regexp.MustCompile("<meta property=.og:image. content=.([^\"']*)")
			//rType         = regexp.MustCompile("REGEX")
			//rPropertyType = regexp.MustCompile("REGEX")
			//rCity         = regexp.MustCompile("REGEX")
			//rRegion       = regexp.MustCompile("REGEX")
			//rPostcode     = regexp.MustCompile("REGEX")
			//rAdress       = regexp.MustCompile("REGEX")
			//rLongitude    = regexp.MustCompile("REGEX")
			//rLatitude     = regexp.MustCompile("REGEX")
		)

		adP := verticals.Homes{
			Url:     ad,
			Title:   rTitle.FindStringSubmatch(content)[1],
			Content: rContent.FindStringSubmatch(content)[1],
			Price: verticals.Price{
				Value:    rprice.FindStringSubmatch(content)[1],
				Currency: verticals.CURRENCY_EUR,
			},
			FloorArea: verticals.FloorArea{
				Value: rFloorArea.FindStringSubmatch(content)[1],
				Unit:  verticals.SQUARED_METERS,
			},
			PlotArea: verticals.PlotArea{
				Value: rPlotArea.FindStringSubmatch(content)[1],
				Unit:  verticals.SQUARED_METERS,
			},
			Rooms:     rRooms.FindStringSubmatch(content)[1],
			Bathrooms: rBathrooms.FindStringSubmatch(content)[1],
			Pictures:  rPictures.FindAllStringSubmatch(content, -1), // -1 means we want All the images, we can limit the number of images
			//Type:         rType.FindStringSubmatch(content)[1],
			//PropertyType: rPropertyType.FindStringSubmatch(content)[1],
			//City:      rCity.FindStringSubmatch(content)[1],
			//Region:    rRegion.FindStringSubmatch(content)[1],
			//Postcode:  rPostcode.FindStringSubmatch(content)[1],
			//Adress:    rAdress.FindStringSubmatch(content)[1],
			//Longitude: rLongitude.FindStringSubmatch(content)[1],
			//Latitude:  rLatitude.FindStringSubmatch(content)[1],
		}

		//Show the info of the ads parsed
		log.Println(adP)
	}
}
