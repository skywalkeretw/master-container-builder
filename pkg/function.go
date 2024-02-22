package pkg

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

func writeFuctionToFile(fd FuncData, dir string) error {
	fmt.Println("Writing function to file: ", fd.Language)
	var file string
	func_call := fmt.Sprintf("%s(%s)", fd.Name, fd.FuncInput)
	switch fd.Language {
	case "javascript":
		replacePlaceholder(filepath.Join(dir, "http.js"), "{{FUNCTION_CALL}}", func_call)
		replacePlaceholder(filepath.Join(dir, "messaging.js"), "{{FUNCTION_CALL}}", func_call)
		file = "function.js"
	case "python":
		replacePlaceholder(filepath.Join(dir, "httpserver.py"), "{{FUNCTION_NAME}}", fd.Name)
		replacePlaceholder(filepath.Join(dir, "messaging.py"), "{{FUNCTION_NAME}}", fd.Name)
		replacePlaceholder(filepath.Join(dir, "httpserver.py"), "{{FUNCTION_CALL}}", func_call)
		replacePlaceholder(filepath.Join(dir, "messaging.py"), "{{FUNCTION_CALL}}", func_call)
		file = "function.py"
	case "golang":
		replacePlaceholder(filepath.Join(dir, "http.go"), "{{FUNCTION_CALL}}", func_call)
		replacePlaceholder(filepath.Join(dir, "messaging.go"), "{{FUNCTION_CALL}}", func_call)
		file = "function.go"
	}
	filePath := filepath.Join(dir, file)
	fmt.Println("Writing function to filePath: ", filePath)

	decodedCodeBytes, err := base64.StdEncoding.DecodeString(fd.Code)
	if err != nil {
		return fmt.Errorf("error decoding: %v", err.Error())
	}
	fmt.Println("Writing function to decodedCodeBytes: ", decodedCodeBytes)

	return writeStringToFile(decodedCodeBytes, filePath)
}

// WriteFunctionsToDisk writes the OpenAPI and AsyncAPI Spec if not empty
func writeSpecToFile(fd FuncData, dir string) error {
	var errMSG error
	if fd.OpenAPISpec != "" {
		filePath := filepath.Join(dir, "openapi.json")
		decodedBytes, err := base64.StdEncoding.DecodeString(fd.OpenAPISpec)
		if err != nil {
			fmt.Println("Error decoding:", err)
			return fmt.Errorf("error decodeing %v", err.Error())
		}

		err = writeStringToFile(decodedBytes, filePath)
		if err != nil {
			errMSG = fmt.Errorf("failed to create openapi.json: %w", err)
		}
	}

	if fd.AsyncAPISpec != "" {
		filePath := filepath.Join(dir, "asyncapi.json")
		decodedBytes, err := base64.StdEncoding.DecodeString(fd.OpenAPISpec)
		if err != nil {
			fmt.Println("Error decoding:", err)
			return fmt.Errorf("error decodeing %v", err.Error())
		}

		err = writeStringToFile(decodedBytes, filePath)
		if err != nil {
			errMSG = fmt.Errorf("%vfailed to create asyncapi.json: %w", errMSG, err)
		}
	}

	return errMSG

}

func replacePlaceholder(filePath, placeholder, replacement string) error {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Perform replacement
	replacedContent := bytes.ReplaceAll(content, []byte(placeholder), []byte(replacement))

	// Write the modified content back to the file
	err = os.WriteFile(filePath, replacedContent, 0644)
	if err != nil {
		return err
	}

	return nil
}
