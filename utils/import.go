package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

func ImportData(path string, structure *Invoice, flags *pflag.FlagSet) error {
	fileText, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to read file")
	}

	var b []byte
	var byteBuffer [][]byte
	flags.Visit(func(f *pflag.Flag) {
		if f.Value.Type() != "string" {
			b = []byte(fmt.Sprintf(`{"%s":%s}`, f.Name, f.Value))
		} else {
			b = []byte(fmt.Sprintf(`{"%s":"%s"}`, f.Name, f.Value))
		}
		byteBuffer = append(byteBuffer, b)
	})

	if strings.HasSuffix(path, ".json") {
		err = importJson(fileText, structure)
	} else if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		err = importYaml(fileText, structure)

	} else {
		return fmt.Errorf("unsupported file type")
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, bytes := range byteBuffer {
		err = importJson(bytes, structure)
		if err != nil {
			log.Fatal(err)
		}
	}

	return err
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
