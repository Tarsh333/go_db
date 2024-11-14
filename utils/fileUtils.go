package utils

// This file will contain all file and folder related operations
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// CreateFolder creates a folder at the specified path if it doesn't exist.
func CreateFolder(path string) error {
	if checkIfFolderExists(path) {
		log.Println(path, "folder already exists")
		return nil
	}
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to create folder %s: %w", path, err)
	}
	log.Println("Folder created:", path)
	return nil
}

// DeleteFolder deletes a folder at the specified path.
func DeleteFolder(path string) error {
	if !checkIfFolderExists(path) {
		log.Println(path, "folder does not exist")
		return nil
	}
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to delete folder %s: %w", path, err)
	}
	log.Println("Folder deleted:", path)
	return nil
}

// AddFile creates a file with the specified data in a given folder.
func AddFile(path, fileName string, data string) error {
	if !checkIfFolderExists(path) {
		return fmt.Errorf("directory %s does not exist", path)
	}
	if checkIfFileExists(path, fileName) {
		log.Println(fileName, "file already exists. overwriting")
	}
	err := os.WriteFile(filepath.Join(path, fileName), []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fileName, err)
	}
	fmt.Println(data)
	log.Println("File created:", fileName)
	return nil
}

func AddJSONFile(path, fileName string, data interface{}) error {
	fileName = fileName + ".json"
	fmt.Println(fileName)
	if !checkIfFolderExists(path) {
		return fmt.Errorf("directory %s does not exist", path)
	}
	if checkIfFileExists(path, fileName) {
		log.Println(fileName, "file already exists")
		return nil
	}

	// Marshal the data to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	// Write the JSON data to file
	err = os.WriteFile(filepath.Join(path, fileName), jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fileName, err)
	}
	log.Println("JSON file created:", fileName)
	return nil
}

// RemoveFile removes a file at the specified path.
func RemoveFile(path, fileName string) error {
	if !checkIfFolderExists(path) {
		return fmt.Errorf("directory %s does not exist", path)
	}
	if !checkIfFileExists(path, fileName) {
		log.Println(fileName, "file does not exist")
		return nil
	}
	err := os.Remove(filepath.Join(path, fileName))
	if err != nil {
		return fmt.Errorf("failed to remove file %s: %w", fileName, err)
	}
	log.Println("File removed:", fileName)
	return nil
}

func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %s", err)
	}
	return string(content), nil
}

func ReadJSONFile(filePath string, v interface{}) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %s", err)
	}

	// Unmarshal the JSON content into the provided struct or map
	err = json.Unmarshal(content, v)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	return nil
}

// checkIfFolderExists checks if a folder exists at the given path.
func checkIfFolderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

// checkIfFileExists checks if a file exists in the given folder.
func checkIfFileExists(path, fileName string) bool {
	fullPath := filepath.Join(path, fileName)
	info, err := os.Stat(fullPath)
	if os.IsNotExist(err) || info.IsDir() {
		return false
	}
	return true
}
func IsValidJSON(data []byte) bool {
	var js interface{}
	return json.Unmarshal(data, &js) == nil
}
func MergeJSONStrings(jsonStrings ...string) (string, error) {
	var mergedData []map[string]interface{}

	for _, jsonString := range jsonStrings {
		// Attempt to parse as an array of objects
		var objArray []map[string]interface{}
		if err := json.Unmarshal([]byte(jsonString), &objArray); err == nil {
			// If successful, add all elements to the mergedData slice
			mergedData = append(mergedData, objArray...)
		} else {
			// If not an array, attempt to parse as a single object
			var obj map[string]interface{}
			if err := json.Unmarshal([]byte(jsonString), &obj); err != nil {
				return "", err
			}
			// Wrap the single object in a slice and append it to mergedData
			mergedData = append(mergedData, obj)
		}
	}

	mergedJSON, err := json.Marshal(mergedData)
	if err != nil {
		return "", err
	}
	return string(mergedJSON), nil
}
