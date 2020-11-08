package main

import (
	"github.com/draco/core"
)

func main() {

	source := core.Source{
		Id:      core.ID,
		Country: core.COUNTRY,
		Url:     core.URL,

		Category: []core.Category{
			{
				Type: core.FOR_SALE,
				Urls: []string{
					"https://www.pisos.com/venta/pisos-a_coruna",
					"https://www.pisos.com/venta/pisos-barcelona",
				},
			},
			{
				Type: core.FOR_RENT,
				Urls: []string{
					"https://www.pisos.com/alquiler/pisos-a_coruna",
					"https://www.pisos.com/alquiler/pisos-barcelona",
				},
			},
		},
	}

	saleUrls, _ := source.ClassifyHomes()

	for _, url := range saleUrls {

		GetAdUrls, count := core.GetAdUrls(url)
		core.ProcessAdList(GetAdUrls, count)

	}

}
