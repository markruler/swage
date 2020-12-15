package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func parse(jsonPath string) (*SwaggerAPI, error) {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Fatalf("%s\n", err)
		return nil, err
	}
	log.Printf("%s\n", "Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("%s\n", err)
		return nil, err
	}

	var swaggerAPI SwaggerAPI
	json.Unmarshal(byteValue, &swaggerAPI)

	return &swaggerAPI, err
}
