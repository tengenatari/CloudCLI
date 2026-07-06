package main

import (
	"CloudCLI/config"
	"CloudCLI/internal/cli"
	"CloudCLI/internal/service"
	"fmt"
	"log"
)

func main() {
	configStruct, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	profileService := service.NewProfileService()
	CLI, err := cli.NewCLI(configStruct, profileService)

	if err != nil {
		log.Fatal(err)
	}

	err = CLI.Run()
	if err != nil {
		fmt.Println(err)
	}
}
