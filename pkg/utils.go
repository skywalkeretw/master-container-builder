package pkg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
)

func GetEnvSting(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	envValue, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	value, err := strconv.Atoi(envValue)

	if err != nil {
		return defaultValue
	}

	return value
}

// GenerateTempFolder generates a temporary folder with a random name using UUID
// and creates missing folders if specified in the path
func GenerateTempFolder() (string, error) {
	// Generate a random UUID for the folder name
	randomFolderName := fmt.Sprintf("build_ctx_%s", uuid.New())

	// Create the full path for the temporary folder
	tempFolderPath := filepath.Join("build", randomFolderName)

	// Create missing folders if specified in the path
	if err := os.MkdirAll(filepath.Dir(tempFolderPath), os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create missing folders: %v", err)
	}

	// Create the temporary folder
	err := os.Mkdir(tempFolderPath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create temporary folder: %v", err)
	}

	return tempFolderPath, nil
}

func CopyFile(src, dst string) error {
	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(filepath.Join(dst, filepath.Base(src)))
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Sync to ensure the data is written to the file system
	err = dstFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
