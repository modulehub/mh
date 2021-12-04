package util

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//CredentialsViper instance of viper
var CredentialsViper *viper.Viper = viper.New()

//InitCredentialsViper returns viper instance
func InitCredentialsViper() {
	log.Info("InitCredentialsViper")
	// stub init
	configHome, err := os.UserHomeDir()
	if err != nil { // handle failed create
		log.Fatal(err)
	} else {
		log.Info("homedir access verified")
	}

	configName := ".mh.credentials"
	configType := "yaml"
	configPath := filepath.Join(configHome, ".modulehub")
	configFile := filepath.Join(configPath, configName+"."+configType)

	err = os.MkdirAll(configPath, 0755)
	if err != nil { // handle failed create
		log.Info(err)
	} else {
		log.Info("config file exists")
	}
	configName = ".mh.credentials"
	_, err = os.OpenFile(configFile, os.O_CREATE, 0644)
	if err != nil { // handle failed create
		log.Info(err)
	} else {
		log.Info("config file exists")
	}

	CredentialsViper.AddConfigPath(configPath)
	CredentialsViper.SetConfigName(configName)
	CredentialsViper.SetConfigType(configType)
}

//GetCredetialsViper instance
func GetCredetialsViper() *viper.Viper {
	log.Info("GetCredetialsViper")

	return CredentialsViper
}
