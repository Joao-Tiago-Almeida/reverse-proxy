package memory

import (
	"reflect"
	"testing"
)

var testData = []map[string]interface{}{
	{"key": "value1"},
	{"key": "value2", "key2": map[string]interface{}{"field2": "value2"}},
	{"key": map[string]interface{}{"field": "value"}},
}

func TestFindOne(t *testing.T) {
	tests := []struct {
		name         string
		desiredKey   string
		desiredValue string
		expected     interface{}
	}{
		{
			name:         "Existing key to string value",
			desiredKey:   "key",
			desiredValue: "value1",
			expected:     map[string]interface{}{"key": "value1"},
		},
		{
			name:         "Existing key to map value",
			desiredKey:   "key",
			desiredValue: "value2",
			expected:     map[string]interface{}{"key": "value2", "key2": map[string]interface{}{"field2": "value2"}},
		},
		{
			name:       "Empty key",
			desiredKey: "",
			expected:   nil,
		},
		{
			name: "panic when memory storage is not initialized",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer Drop()
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("TestFindOne(%s) panicked unexpectedly", test.name)
				}
			}()

			var m memoryDB = New()
			// Insert some test data
			for _, data := range testData {
				m.Insert(data)
			}

			result, _ := m.FindOne(test.desiredKey, test.desiredValue)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("TestFindOne(%s): have %v; want %v", test.name, result, test.expected)
			}
		})
	}
}

func TestFind(t *testing.T) {
	defer Drop()
	var m memoryDB
	New()

	// Insert some test data
	for _, data := range testData {
		m.Insert(data)
	}

	tests := []struct {
		name     string
		filters  map[string]string
		expected []interface{}
	}{
		{
			name:     "Existing keys",
			filters:  map[string]string{"key": "value1"},
			expected: []interface{}{map[string]interface{}{"key": "value1"}},
		},
		{
			name:     "Non-existing keys",
			filters:  map[string]string{"": "value1"},
			expected: nil,
		},
		{
			name:     "Non-existing values",
			filters:  map[string]string{"key": ""},
			expected: nil,
		},
		{
			name:     "Existing and non-existing keys",
			filters:  map[string]string{"key": "value1", "": "value1"},
			expected: []interface{}{map[string]interface{}{"key": "value1"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer Drop()
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("TestFindOne(%s) panicked unexpectedly", test.name)
				}
			}()

			New()
			// Insert some test data
			for _, data := range testData {
				m.Insert(data)
			}

			result, _ := m.Find(test.filters)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("TestFind(%s): have %v; want %v", test.name, result, test.expected)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	defer Drop()

	var m memoryDB
	New()

	// Insert some test data
	for _, data := range testData {
		m.Insert(data)
	}

	tests := []struct {
		name         string
		desiredKey   string
		desiredValue string
	}{
		{
			name:         "Delete existing key",
			desiredKey:   "key",
			desiredValue: "value1",
		},
		{
			name: "Delete non-existing key",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("TestDeleteOne(%s) panicked unexpectedly", test.name)
				}
			}()

			m.Delete(test.desiredKey, test.desiredValue)
			result, _ := m.FindOne(test.desiredKey, test.desiredValue)
			if result != nil {
				t.Errorf("TestDeleteOne(%s): key %s was not deleted", test.name, test.desiredKey)
			}
		})
	}
}

func Test_checkIfExsits(t *testing.T) {
	var m memoryDB
	tests := []struct {
		name string
		f    func()
	}{
		{
			name: "Insert",
			f: func() {
				m.Insert(nil)
			},
		},
		{
			name: "Find",
			f: func() {
				m.Find(nil)
			},
		},
		{
			name: "FindOne",
			f: func() {
				m.FindOne("", "")
			},
		},
		{
			name: "Delete",
			f: func() {
				m.Delete("", "")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					t.Errorf("Test_checkIfExsits(%s) didn't panicked", test.name)
				}
			}()

			test.f()
		})
	}
}
