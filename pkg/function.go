package pkg

import (
	"fmt"
	"path/filepath"
)

func writeFuctionToFile(fd FuncData, dir string) error {
	var file string
	switch fd.Language {
	case "javascript":
		file = "function.js"
	case "python":
		file = "function.py"
	case "golang":
		file = "function.go"
	}
	filePath := filepath.Join(dir, file)

	return writeStringToFile(fd.Code, filePath)
}

// WriteFunctionsToDisk writes the OpenAPI and AsyncAPI Spec if not empty
func writeSpecToFile(fd FuncData, dir string) error {
	var errMSG error
	if fd.OpenAPISpec != "" {
		filePath := filepath.Join(dir, "openapi.json")
		err := writeStringToFile(fd.OpenAPISpec, filePath)
		if err != nil {
			errMSG = fmt.Errorf("failed to create openapi.json: %w", err)
		}
	}

	if fd.AsyncAPISpec != "" {
		filePath := filepath.Join(dir, "asyncapi.json")
		err := writeStringToFile(fd.OpenAPISpec, filePath)
		if err != nil {
			errMSG = fmt.Errorf("%vfailed to create asyncapi.json: %w", errMSG, err)
		}
	}

	return errMSG

}
