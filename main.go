package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Health struct {
	Workouts struct {
		name string `json:"name"`

	} `json:"workouts"`
}

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
			// filePath := dirPath + "/" + file.Name()
			filePath := "/Users/saavedj/Library/Mobile Documents/com~apple~CloudDocs/health-data/HealthAutoExport-2024-10-21-2025-01-19.json"

			content, err := os.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

			var jsonData map[string]interface{}
			err = json.Unmarshal(content, &jsonData)
			if err != nil {
				panic(err)
			}

			fmt.Println("Fields (keys):")
			for key := range jsonData {
				fmt.Println(key)
			}


		}
	}

}