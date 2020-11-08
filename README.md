## Crawler for Real estate!

It's a library that allows you to crawl through real estate websites [any] and get all the content from there.

## Installation

Just run this command and make sure you import it into your project.

```bash
go get github.com/Neufal777/crawler
```

## Usage

```golang
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

```

### Result

Return a Struct(Homes struct) with all the information of each AD parsed
```bash
{ 
     https://www.pisos.com/comprar/duplex-bertamirans-5911333496_995784/
     Dúplex en venta en Avenida de la Mahía, cerca de Calle de las Pesqueiras 
	 Se vende piso en el centro de ,  a lado del supermercado y mercadona.
	 Reformado, con ascensor, estancias amplias y luminosas.
	 Listo para entrar. Planta baja: cocina amueblad...   
     {119.000 € EUR}
     {120 m2}
     {120 m2}
     3 habs
     2 baños 
     [ https://fotos.imghs.net/xl/995784/5911333496.995784/995784_5911333496_18_20200716200528618.jpg]
}
```
