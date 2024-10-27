package memory

var db []map[string]interface{}

type memoryDB struct{}

func New() memoryDB {
	// Initialize the memory Storage
	db = make([]map[string]interface{}, 0)

	return memoryDB{}
}

func (m memoryDB) Insert(data map[string]interface{}) error {
	m.checkIfExsits()

	// Append the data to the memory storage
	db = append(db, data)

	return nil
}

func (m memoryDB) Find(desiredValues []string) (findings []interface{}, err error) {
	m.checkIfExsits()

	// For each desired key, find the data in the memory storage
	for _, desiredKey := range desiredValues {

		value, err := m.FindOne(desiredKey)
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

func (m memoryDB) FindOne(desiredKey string) (finding interface{}, err error) {
	m.checkIfExsits()

	// Look for the desired key in the memory storage
	for _, dict := range db { // Loop through the memory storage
		for key, value := range dict { // Loop through the dictionary
			if desiredKey == key {
				return value, nil
			}
		}
	}

	return nil, nil
}

func (m memoryDB) DeleteOne(desiredKey string) error {
	m.checkIfExsits()

	// Look for the desired key in the memory storage
	for index, dict := range db {
		for key, _ := range dict {
			if desiredKey == key {
				// Remove the data from the memory storage
				db = append(db[:index], db[index+1:]...)
			}
		}
	}
	return nil
}

func Drop() {
	db = nil
}

func (m memoryDB) checkIfExsits() {
	if db == nil {
		panic("Memory storage is not initialized")
	}
}
