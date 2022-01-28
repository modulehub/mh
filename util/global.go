package util

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
)

//GetURL return api url
func GetURL() (*url.URL, error) {
	return url.Parse(os.Getenv("MH_API_BASE_URL"))
}

//GetEnv loads value from env if exists with fallback to a provided default value
func GetEnv(name string, defVal string) string {
	if val, exists := os.LookupEnv(name); exists {
		return val
	}

	return defVal
}

//OpenURL opens a given url
func OpenURL(url string) error {
	var err error
	if url != "" {
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}
		return err
	}
	return errors.New("please provide a valid url")
}
