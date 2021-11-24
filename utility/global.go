package utility

import (
	"net/url"

	"github.com/spf13/viper"
)

//GetURL return api url
func GetURL() url.URL {
	var baseURL = url.URL{
		Scheme: viper.GetString("api_scheme"),
		Host:   viper.GetString("api_host"),
		Path:   "/api",
	}
	return baseURL
}
