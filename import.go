package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func importData(path string, structure *Invoice) error {
	fileText, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to read file")
	}

	if strings.HasSuffix(path, ".json") {
		return importJson(fileText, structure)
	}

	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		return importYaml(fileText, structure)
	}

	return fmt.Errorf("unsupported file type")
}

func importJson(text []byte, structure *Invoice) error {
	if !json.Valid(text) {
		return fmt.Errorf("json file not correctly formatted")
	}

	err := json.Unmarshal(text, structure)
	if err != nil {
		return fmt.Errorf("json file not correctly formatted")
	}

	return nil
}

func importYaml(text []byte, structure *Invoice) error {
	err := yaml.Unmarshal(text, structure)
	if err != nil {
		return fmt.Errorf("yaml file not correctly formatted")
	}

	return nil
}
