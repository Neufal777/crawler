package main

import (
	"log"

	"github.com/draco/core"
)

func main() {

	fuente := core.Source{
		Id:       core.ID,
		Country:  core.COUNTRY,
		Vertical: core.VERTICAL,
		Url:      core.URL,

		Category: []core.Category{
			{
				Type: core.FOR_SALE,
				Urls: []string{
					"https://www.pisos.com/venta/pisos-a_coruna",
					"https://www.pisos.com/venta/pisos-barcelona",
					"https://www.pisos.com/venta/pisos-girona",
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

	//log.Println(fuente)

	sale, rent := fuente.ClassifyCat()

	log.Println(sale)
	log.Println(rent)
	//core.Pagination(rent)
	//core.Pagination(sale)

}
