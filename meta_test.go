package bio_test

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Test functions should be housed in packages following the "_test" convention
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
			if strings.HasPrefix(scan.Text(), "package ") && !strings.HasSuffix(scan.Text(), "_test") {
				t.Errorf("Test file %q, is not in a test pkg. Got: %s", path, scan.Text())
			}
		}
		return nil
	})
}

// Tests, Examples, and Benchmarks should be sectioned together, not intermixed
func TestOrder_Test_Example_Benchmark(t *testing.T) {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), "_test.go") {
			visited := make(map[string]struct{})
			file, err := os.Open(path)
			defer file.Close()
			if err != nil {
				t.Errorf("Could not open: %q", path)
			}
			scan := bufio.NewScanner(file)
			line := 1
			for scan.Scan() {
				switch {
				case strings.Contains(scan.Text(), "func Test"):
					visited["test"] = struct{}{}
					if _, err := visited["example"]; err {
						t.Errorf("\nExamples found above test on line %d in %s", line, file.Name())
					}
					if _, err := visited["benchmark"]; err {
						t.Errorf("\nBenchmarks found above test on line %d in %s", line, file.Name())
					}
				case strings.Contains(scan.Text(), "func Example"):
					visited["example"] = struct{}{}
					if _, err := visited["benchmark"]; err {
						t.Errorf("\nBenchmark found above example on line %d in %s", line, file.Name())
					}
				case strings.Contains(scan.Text(), "func Benchmark"):
					visited["benchmark"] = struct{}{}
				default:
				}
				line++
			}
		}
		return nil
	})
}
