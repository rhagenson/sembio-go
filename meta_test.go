package bio_test

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTestsAreInSeparatePkg(t *testing.T) {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), "_test.go") {
			file, err := os.Open(path)
			defer file.Close()
			if err != nil {
				t.Errorf("Could not open: %q", path)
			}
			scan := bufio.NewScanner(file)
			scan.Scan()
			if !strings.HasSuffix(scan.Text(), "_test") {
				t.Errorf("Test file %q, is not in a test pkg. Got: %s", path, scan.Text())
			}
		}
		return nil
	})
}
