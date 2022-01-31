/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/modulehub/mh/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	var err error
	err = os.Setenv("MH_API_BASE_URL", "https://api.v2.modulehub.io/")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("MH_APP_BASE_URL", "https://app.modulehub.io/")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("MH_APP_TERRAFORM_REGISTRY_URL", "https://registry.v2.modulehub.io/")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("MH_APP_HELM_REGISTRY_URL", "https://registry.modulehub.io/")
	if err != nil {
		panic(err)
	}

	err = godotenv.Overload()
	if err != nil {
		log.Info("continue without .env file")
	}
	cmd.Execute()
}
