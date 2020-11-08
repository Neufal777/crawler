package core

import (
	"io/ioutil"
	"net/http"
)

func DownloadHtml(url string) string {

	/*
		This function downloads an html source code and
		turns it into a []byte
	*/

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	h, err := ioutil.ReadAll(resp.Body)

	html := string(h[:])
	//htmlString := string(html[:])

	return html
}
