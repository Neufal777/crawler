package core

import (
	"io/ioutil"
	"net/http"
)

func DownloadHtml(url string) []byte {

	/*
		This function downloads an html source code and
		turns it into a []byte
	*/

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)

	//htmlString := string(html[:])

	return html
}
