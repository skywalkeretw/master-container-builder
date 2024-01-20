package pkg

import "path/filepath"

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
