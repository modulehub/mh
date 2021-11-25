package util

import (
	"net/url"
	"os"

	"github.com/spf13/viper"
)

//GetURL return api url
func GetURL() (*url.URL, error) {
	return url.Parse(viper.GetString("api_url"))
}

//GetEnv loads value from env if exists with fallback to a provided default value
func GetEnv(name string, defVal string) string {
	if val, exists := os.LookupEnv(name); exists {
		return val
	}

	return defVal
}
