package tools

import (
	"encoding/json"
	"fmt"
	"xs/pkg/io"
)

import "github.com/xeipuuv/gojsonschema"

// go get -u github.com/go-bindata/go-bindata/...

func jsonLoadAndValidate(data string, schema string) {
	// Load the schema
	schemaLoader := gojsonschema.NewStringLoader(schema)

	// Load the data
	dataLoader := gojsonschema.NewStringLoader(data)

	// Perform validation
	result, err := gojsonschema.Validate(schemaLoader, dataLoader)
	if err != nil {
		io.Fatal(err)
	}

	// Check if the data is valid
	if result.Valid() {
		io.Println("The JSON data is valid.")
	} else {
		io.Println("The JSON data is not valid. Validation errors:")
		for _, err := range result.Errors() {
			fmt.Printf("- %s\n", err)
		}
	}
}

func ValidateJson(jsonFile string, st any) {
	bytesFromFile, err := ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(bytesFromFile), st)
	if err != nil {
		panic(err)
	}
}