package external

import (
	"fmt"
	"log"
)

type storage struct {
	FilePath string
}

type Record interface {
	GetID() string
	GenerateUUID()
}

type Storage interface {
	FindAll() ([]Record, error)
	FindOne(id string) (Record, error)
	Insert(record Record) (Record, error)
	Update(record Record) (Record, error)
	Delete(record Record) error
}

func NewJsonStorage(filePath string) Storage {
	storage := &storage{FilePath: filePath}

	isExisted := IsFileExists(filePath)

	if isExisted {
		log.Printf("DB file found: %s", filePath)
		return storage
	}

	// create new file if not exists
	err := CreateJSONFile(filePath)
	if err != nil {
		log.Printf("error creating file: %v", err)
		return nil
	}

	log.Printf("DB file created: %s", filePath)

	return storage
}

func (s storage) FindAll() ([]Record, error) {
	json, err := ReadJSON(s.FilePath)
	if err != nil {
		return nil, err
	}

	records := make([]Record, 0)
	for _, v := range json {
		records = append(records, v)
	}

	return records, nil
}

func (s storage) FindOne(id string) (Record, error) {
	records, err := s.FindAll()
	if err != nil {
		return nil, err
	}

	for _, v := range records {
		if v.GetID() == id {
			return v, nil
		}
	}

	return nil, fmt.Errorf("record not found")
}

func (s storage) Insert(record Record) (Record, error) {
	records, err := s.FindAll()
	if err != nil {
		return nil, err
	}

	if record.GetID() == "" {
		record.GenerateUUID()
	}

	records = append(records, record)

	err = WriteJSON(s.FilePath, records)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (s storage) Update(record Record) (Record, error) {
	_, err := s.FindOne(record.GetID())
	if err != nil {
		return nil, fmt.Errorf("record not found")
	}

	records, err := s.FindAll()
	if err != nil {
		return nil, err
	}

	for i, v := range records {
		if v.GetID() == record.GetID() {
			records[i] = record
			break
		}
	}

	err = WriteJSON(s.FilePath, records)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (s storage) Delete(record Record) error {
	if record.GetID() == "" {
		return fmt.Errorf("field ID is required")
	}

	records, err := s.FindAll()
	if err != nil {
		return err
	}

	for i, v := range records {
		if v.GetID() == record.GetID() {
			records = append(records[:i], records[i+1:]...)
			break
		}
	}

	err = WriteJSON(s.FilePath, records)
	if err != nil {
		return err
	}

	return nil
}
