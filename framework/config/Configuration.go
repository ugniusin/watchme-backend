package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "./config.json"

type Configuration struct {
	Instagram map[string]string
	Database map[string]string
}

func NewConfiguration() *Configuration {
	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Filename: " + fileName)

	return &configuration
}
