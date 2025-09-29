package lang

import (
	"reflect"
	"testing"
)

func TestParseRequestedLanguages(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []Language
		expectError bool
	}{
		{
			name:        "parse 'all' returns all supported languages",
			input:       "all",
			expected:    []Language{dotNet, goLang, java, nodeJS, python},
			expectError: false,
		},
		{
			name:        "parse 'ALL' (case insensitive) returns all supported languages",
			input:       "ALL",
			expected:    []Language{dotNet, goLang, java, nodeJS, python},
			expectError: false,
		},
		{
			name:        "parse empty string returns empty slice",
			input:       "",
			expected:    []Language{},
			expectError: true,
		},
		{
			name:        "parse single valid language",
			input:       "go",
			expected:    []Language{goLang},
			expectError: false,
		},
		{
			name:        "parse multiple valid languages",
			input:       "go,python,java",
			expected:    []Language{goLang, python, java},
			expectError: false,
		},
		{
			name:        "parse all languages explicitly",
			input:       "nodejs,go,python,dotnet,java",
			expected:    []Language{nodeJS, goLang, python, dotNet, java},
			expectError: false,
		},
		{
			name:        "parse with empty token should fail",
			input:       "go,,python",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "parse invalid language returns error",
			input:       "invalid",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "parse mixed valid and invalid languages returns error",
			input:       "go,THIS ISNT A SUPPORTED LANGUAGE,python",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "parse languages with spaces returns error",
			input:       "go, python, java",
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseRequestedLanguages(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestParseRequestedLanguagesEmptyString(t *testing.T) {
	result, err := ParseRequestedLanguages("")

	if err == nil {
		t.Error("expected error for empty string but got none")
	}

	if result != nil {
		t.Errorf("expected nil result for empty string but got %v", result)
	}
}

func TestParseRequestedLanguagesGroupOfLanguages(t *testing.T) {
	input := "python,nodejs,dotnet"
	expected := []Language{python, nodeJS, dotNet}

	result, err := ParseRequestedLanguages(input)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestParseRequestedLanguagesAll(t *testing.T) {
	result, err := ParseRequestedLanguages("all")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := AllSupportedLanguages()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
