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
	// Get absolute path
	absPath, err := filepath.Abs(tempFolderPath)
	if err != nil {
		return "", err
	}

	return absPath, nil
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

func CopyFolder(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			// Create the destination directory
			return os.MkdirAll(destPath, os.ModePerm)
		}

		// Open the source file
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		// Create the destination file
		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		// Copy the contents from source to destination
		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}

		// Sync to ensure the data is written to the file system
		err = destFile.Sync()
		if err != nil {
			return err
		}

		return nil
	})
}

func writeStringToFile(content, filePath string) error {
	// Write the string content to the file
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

func getCurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}
