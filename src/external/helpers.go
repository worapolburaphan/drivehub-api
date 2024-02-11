package external

import (
	"encoding/json"
	"fmt"
	"github.com/drivehub-api/src/entity"
	"io"
	"os"
	"path/filepath"
)

func IsFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return true
}

func CreateJSONFile(filePath string) error {
	dirName := filepath.Dir(filePath)
	// check if file exists
	_, err := os.Stat(filePath)
	if err == nil {
		// ignore create file
		return nil
	} else {
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString("[]")
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

func ReadJSON(filename string) ([]entity.JSONEntity, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var records []entity.JSONEntity
	err = json.Unmarshal(data, &records)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return records, nil
}

func WriteJSON(filename string, data []Record) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	encoded, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	_, err = file.Write(encoded)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}
