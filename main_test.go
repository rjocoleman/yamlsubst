package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestProcessTemplate(t *testing.T) {
	tests := []struct {
		name           string
		yamlContent    string
		templateInput  string
		expectedOutput string
		expectError    bool
	}{
		{
			name: "basic template",
			yamlContent: `
name: John
age: 30
hobbies:
  - reading
  - writing
`,
			templateInput:  "Name: {{.name}}, Age: {{.age}}",
			expectedOutput: "Name: John, Age: 30",
			expectError:    false,
		},
		{
			name: "using custom functions",
			yamlContent: `
items:
  - apple
  - banana
text: HELLO
`,
			templateInput:  `{{join .items ","}} {{toLower .text}}`,
			expectedOutput: "apple,banana hello",
			expectError:    false,
		},
		{
			name:           "invalid yaml",
			yamlContent:    "invalid: [yaml: content",
			templateInput:  "shouldn't matter",
			expectedOutput: "",
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary YAML file
			tmpfile, err := os.CreateTemp("", "test-*.yaml")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write([]byte(tt.yamlContent)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			// Create input pipe for template content
			inR, inW := io.Pipe()
			go func() {
				_, err := inW.Write([]byte(tt.templateInput))
				if err != nil {
					t.Error(err)
				}
				inW.Close()
			}()

			// Create output buffer
			var output bytes.Buffer

			// Process template
			err = processTemplate(tmpfile.Name(), inR, &output)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("processTemplate() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !tt.expectError && strings.TrimSpace(output.String()) != strings.TrimSpace(tt.expectedOutput) {
				t.Errorf("processTemplate() = %v, want %v", output.String(), tt.expectedOutput)
			}
		})
	}
}

func TestInterfaceJoiner(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		sep      string
		expected string
	}{
		{
			name:     "basic join",
			input:    []interface{}{"a", "b", "c"},
			sep:      ",",
			expected: "a,b,c",
		},
		{
			name:     "empty slice",
			input:    []interface{}{},
			sep:      ",",
			expected: "",
		},
		{
			name:     "single item",
			input:    []interface{}{"a"},
			sep:      ",",
			expected: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := interfaceJoiner(tt.input, tt.sep)
			if result != tt.expected {
				t.Errorf("interfaceJoiner() = %v, want %v", result, tt.expected)
			}
		})
	}
}
