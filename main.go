package main

import (
	"log"

	"github.com/draco/core"
)

func main() {

	fuente := core.Source{
		Id:       281,
		Country:  "es",
		Vertical: 2,
		Url:      "https://www.pisos.com/",
		Category: []core.Category{
			{
				Type: "FOR_SALE",
				Urls: []string{
					"https://www.pisos.com/venta/pisos-a_coruna/",
					"https://www.pisos.com/venta/pisos-barcelona/",
					"https://www.pisos.com/venta/pisos-barcelona/",
				},
			},
			{
				Type: "FOR_RENT",
				Urls: []string{
					"https://www.pisos.com/alquiler/pisos-a_coruna/",
					"https://www.pisos.com/alquiler/pisos-barcelona/",
					"https://www.pisos.com/alquiler/pisos-barcelona/",
				},
			},
		},
	}

	//log.Println(fuente)

	_, rent := core.ClassifyCat(&fuente)

	//log.Println(sale)
	log.Println(rent)
	//core.Pagination(cat)
}
