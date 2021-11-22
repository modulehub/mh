package utility

import "net/url"

var baseURL = url.URL{
	Scheme: "http",
	Host:   "localhost:81",
	Path:   "/api",
}

//GetUrl return api url
func GetURL() url.URL {
	return baseURL
}
