package memory

import (
	"reflect"
	"testing"
)

var testData = []map[string]interface{}{
	{"key1": "value1"},
	{"key2": map[string]interface{}{"field2": "value2"}},
}

func TestFindOne(t *testing.T) {
	tests := []struct {
		name       string
		desiredKey string
		expected   interface{}
	}{
		{
			name:       "Existing key to string value",
			desiredKey: "key1",
			expected:   "value1",
		},
		{
			name:       "Existing key to map value",
			desiredKey: "key2",
			expected:   map[string]interface{}{"field2": "value2"},
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

			result, _ := m.FindOne(test.desiredKey)
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
		name          string
		desiredValues []string
		expected      []interface{}
	}{
		{
			name:          "Existing keys",
			desiredValues: []string{"key1", "key2"},
			expected:      []interface{}{"value1", map[string]interface{}{"field2": "value2"}},
		},
		{
			name:          "Non-existing keys",
			desiredValues: []string{"key3", "key4"},
			expected:      nil,
		},
		{
			name:          "Existing and non-existing keys",
			desiredValues: []string{"key1", "key3"},
			expected:      []interface{}{"value1"},
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

			result, _ := m.Find(test.desiredValues)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("TestFind(%s): have %v; want %v", test.name, result, test.expected)
			}
		})
	}
}

func TestDeleteOne(t *testing.T) {
	defer Drop()

	var m memoryDB
	New()

	// Insert some test data
	for _, data := range testData {
		m.Insert(data)
	}

	tests := []struct {
		name       string
		desiredKey string
	}{
		{
			name:       "Delete existing key",
			desiredKey: "key1",
		},
		{
			name:       "Delete non-existing key",
			desiredKey: "key3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("TestDeleteOne(%s) panicked unexpectedly", test.name)
				}
			}()

			m.DeleteOne(test.desiredKey)
			result, _ := m.FindOne(test.desiredKey)
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
				m.FindOne("")
			},
		},
		{
			name: "DeleteOne",
			f: func() {
				m.DeleteOne("")
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
