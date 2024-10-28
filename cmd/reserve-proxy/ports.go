package main

import "github.com/Joao-Tiago-Almeida/reverse-proxy/pkg/database/memory"

type database interface {
	Insert(map[string]interface{}) error
	Find(map[string]string) ([]interface{}, error)
	FindOne(string, string) (interface{}, error)
	Delete(string, string) error
}

func initDatabase() database {
	// Initialize the database
	// Using an in-memory database as a backup or for testing
	// TODO: Implement a file system database, like csv that is contolled by the env variable DATABASE_TYPE OR DATABASE_CSV_PATH
	return memory.New()
}
