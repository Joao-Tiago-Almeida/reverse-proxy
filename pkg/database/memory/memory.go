package memory

import (
	"fmt"
	"os"
)

var db []map[string]interface{}

type memoryDB struct{}

func New() memoryDB {
	// Initialize the memory Storage
	db = make([]map[string]interface{}, 0)

	// Fill the memory storage with some data if spceified in the nev
	if os.Getenv("DEBUG") == "true" {
		fill()
	}

	return memoryDB{}
}

func (m memoryDB) Insert(data map[string]interface{}) error {
	m.checkIfExsits()

	// Append the data to the memory storage
	db = append(db, data)

	return nil
}

func (m memoryDB) Find(filters map[string]string) (findings []interface{}, err error) {
	m.checkIfExsits()

	// For each desired key, find the data in the memory storage
	for desiredKey, desiredValue := range filters {

		value, err := m.FindOne(desiredKey, desiredValue)
		if err != nil {
			return nil, err
		}

		if value != nil {
			findings = append(findings, value)
		}
	}

	if len(findings) == 0 {
		return nil, nil
	}
	return findings, nil
}

func (m memoryDB) FindOne(key string, value string) (finding interface{}, err error) {
	m.checkIfExsits()

	// Look for the desired key in the memory storage
	for _, dBdict := range db { // Loop through the memory storage
		for dBkey, dBvalue := range dBdict { // Loop through the dictionary
			if key == dBkey && value == dBvalue {
				return dBdict, nil
			}
		}
	}

	return nil, nil
}

func (m memoryDB) Delete(key string, value string) error {
	m.checkIfExsits()

	// Look for the desired key in the memory storage
	for index, dict := range db {
		for dBkey, dBvalue := range dict {
			if key == dBkey && value == dBvalue {
				// Remove the data from the memory storage
				db = append(db[:index], db[index+1:]...)
			}
		}
	}
	return nil
}

func (m memoryDB) checkIfExsits() {
	if db == nil {
		panic("Memory storage is not initialized")
	}
}

func Drop() {
	db = nil
}

func findAll() (dbCopy []map[string]interface{}) {
	dbCopy = make([]map[string]interface{}, len(db))
	copy(dbCopy, db)
	return dbCopy
}

func fill() {
	// Data to be filled in the memory storage
	data := []map[string]interface{}{
		{"host": "tarkin", "alias": "localhoss:8002", "createdAt": "2024-10-28T09:00:00Z"},
		{"host": "iris", "alias": "localhoss:8080", "createdAt": "2024-10-28T09:02:47Z"},
	}

	db = append(db, data...)
	fmt.Println(db)
}
