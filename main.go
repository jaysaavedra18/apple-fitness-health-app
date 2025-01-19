package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the directory (iCloud Drive)
	dirPath := "/Users/saavedj/Library/Mobile Documents/com~apple~CloudDocs/health-data"

	files, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	// Sift through the files for json files
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			filePath := dirPath + "/" + file.Name()

			content, err := os.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			fmt.Printf("File: %s\n", file.Name())
			fmt.Println(string(content))
		}
	}


}