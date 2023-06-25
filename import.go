package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func importData(path string, structure *invoiceData) error {
	fileText, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	length := len(path)
	if path[length-5:] == ".json" {
		return importJson(fileText, structure)
	}
	if path[length-5:] == ".yaml" || path[length-4:] == ".yml" {
		return importYaml(fileText, structure)
	}

	return fmt.Errorf("unsupported file type")
}

func importJson(text []byte, structure *invoiceData) error {

	if !json.Valid(text) {
		return fmt.Errorf("json file not correctly formatted")
	}

	err := json.Unmarshal(text, structure)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func importYaml(text []byte, structure *invoiceData) error {

	err := yaml.Unmarshal(text, structure)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
