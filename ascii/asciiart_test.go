package ascii

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetBanner(t *testing.T) {
	// 1. Setup a temporary directory for banners to keep the test environment isolated
	testDir := "banners"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temporary banners directory: %v", err)
	}
	// Clean up the directory entirely after this test finishes
	defer os.RemoveAll(testDir)

	mockData := "line1\r\nline2\r\nline3"
	filename := filepath.Join(testDir, "test_banner.txt")
	err = os.WriteFile(filename, []byte(mockData), 0644)
	if err != nil {
		t.Fatalf("Failed to create test banner file: %v", err)
	}

	// 2. Execute target function
	lines, err := GetBanner(filename)
	if err != nil {
		t.Errorf("GetBanner returned an unexpected error: %v", err)
	}

	// 3. Assert outputs
	if len(lines) != 3 || lines[0] != "line1" {
		t.Errorf("GetBanner did not correctly normalize or split lines. Got: %v", lines)
	}
}

func TestGenerateAscii(t *testing.T) {
	// 1. Setup mock banners directory
	testDir := "banners"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temporary banners directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	// Create a minimal mock banner representing:
	// Space character (ASCII 32) -> 9 lines (1 empty newline + 8 rows of space strings)
	mockBanner := "\n \n \n \n \n \n \n \n \n"
	
	// Pad the banner up to character 'A' (ASCII 65) with placeholder lines
	for i := 33; i < 65; i++ {
		mockBanner += "\n.\n.\n.\n.\n.\n.\n.\n.\n"
	}
	
	// 'A' (ASCII 65) -> 1 empty newline + 8 rows of pattern blocks
	mockBanner += "\nAAA\nAAA\nAAA\nAAA\nAAA\nAAA\nAAA\nAAA\n"

	mockStyleName := "mock_style"
	err = os.WriteFile(filepath.Join(testDir, mockStyleName+".txt"), []byte(mockBanner), 0644)
	if err != nil {
		t.Fatalf("Failed to write mock style asset: %v", err)
	}

	// 2. Table-driven test vectors evaluating your specific processing rules
	tests := []struct {
		name        string
		input       string
		banner      string
		expectError bool
		containStr  string
	}{
		{
			name:        "Valid Single Character Rendering",
			input:       "A",
			banner:      mockStyleName,
			expectError: false,
			containStr:  "AAA",
		},
		{
			name:        "Handling Missing Banner Assets",
			input:       "A",
			banner:      "missing_banner_file",
			expectError: true,
			containStr:  "",
		},
		{
			name:        "Verifying Literal Backslash-n Handling",
			input:       "A\\nA",
			banner:      mockStyleName,
			expectError: false,
			containStr:  "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GenerateAscii(tt.input, tt.banner)
			
			// Assert Error expectations
			if (err != nil) != tt.expectError {
				t.Fatalf("GenerateAscii error status mismatch: got error=%v, expected error status=%v", err, tt.expectError)
			}

			// Assert Output expectations if no error occurred
			if !tt.expectError && tt.containStr != "" && !strings.Contains(res, tt.containStr) {
				t.Errorf("Expected output block to contain %q, but got:\n%s", tt.containStr, res)
			}
		})
	}
}