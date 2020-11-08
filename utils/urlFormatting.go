package utils

func CompleteUrl(parts ...string) string {

	var formattedUrl string

	for _, part := range parts {

		formattedUrl += part
	}
	return formattedUrl
}
