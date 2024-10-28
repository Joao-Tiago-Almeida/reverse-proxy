package utils

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Name  string `omitempty:"true"`
	Age   int
	Email string `json:"mail"`
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]interface{}
	}{
		{
			name:  "Valid struct",
			input: TestStruct{Name: "John", Age: 30, Email: "john@example.com"},
			expected: map[string]interface{}{
				"Name": "John",
				"Age":  float64(30),
				"mail": "john@example.com",
			},
		},
		{
			name:     "Invalid struct",
			input:    nil,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Map(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("%s: Map(%v): have %v; want %v", test.name, test.input, result, test.expected)
			}
		})
	}
}

type ValidTestStruct struct {
	Name  string `omitempty:"true"`
	Age   int
	Email string `json:"mail"`
}

func (v ValidTestStruct) validate() bool {
	return v.Name != "" && v.Age > 0 && v.Email != ""
}

func TestValidMap(t *testing.T) {
	tests := []struct {
		name      string
		input     Validatable
		expected  map[string]interface{}
		wantPanic bool
	}{
		{
			name:  "Valid struct",
			input: ValidTestStruct{Name: "John", Age: 30, Email: "john@example.com"},
			expected: map[string]interface{}{
				"Name": "John",
				"Age":  float64(30),
				"mail": "john@example.com",
			},
			wantPanic: false,
		},
		{
			name:      "Invalid struct",
			input:     ValidTestStruct{Name: "", Age: 0, Email: ""},
			expected:  nil,
			wantPanic: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			func() {
				defer func() {
					if r := recover(); r != nil {
						if !test.wantPanic {
							t.Errorf("%s: ValidMap(%v) panicked unexpectedly", test.name, test.input)
						}
					} else {
						if test.wantPanic {
							t.Errorf("%s: ValidMap(%v) did not panic as expected", test.name, test.input)
						}
					}
				}()

				result := ValidMap(test.input)
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("%s: ValidMap(%v): have %v; want %v", test.name, test.input, result, test.expected)
				}
			}()
		})
	}
}
