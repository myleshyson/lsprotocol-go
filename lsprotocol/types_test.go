package lsprotocol

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestType(t *testing.T) {

	cwd, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	files, err := os.ReadDir(cwd + "/testdata")

	if err != nil {
		t.Error(err)
	}
	for _, file := range files {
		filepath := cwd + "/testdata/" + file.Name()
		if !strings.Contains(file.Name(), "json") {
			continue
		}

		parts := strings.Split(file.Name(), "-")

		testType := parts[0]
		expectedResult := parts[1]
		content, err := os.ReadFile(filepath)

		if err != nil {
			t.Error(err)
		}

		switch testType {
		case "ApplyWorkspaceEditRequest":
			var request ApplyWorkspaceEditRequest

			err := json.Unmarshal(content, &request)

			if err != nil && expectedResult == "True" {
				t.Fatalf("Expected pass in file: %s, got error: %s", filepath, err)
			}

			if err == nil && expectedResult == "False" {
				t.Fatalf("Expected fail in file: %s", filepath)
			}
		case "ApplyWorkspaceEditResponse":
			var request ApplyWorkspaceEditResponse

			err := json.Unmarshal(content, &request)

			if err != nil && expectedResult == "True" {
				t.Fatalf("Expected pass in file: %s, got error: %s", filepath, err)
			}

			if err == nil && expectedResult == "False" {
				t.Fatalf("Expected fail in file: %s", filepath)
			}
		default:
			break
		}
	}
}
