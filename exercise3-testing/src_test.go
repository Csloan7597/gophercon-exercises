package exercise3_testing

import (
	"fmt"
	"strings"
	"testing"
)

// Write a test for strings.HasPrefix
// https://golang.org/pkg/strings/#HasPrefix
// Given the value "main.go", test that it has the prefix "main"
// Remember that your test has to start with the name `Test` and be in an `_test.go` file

// Write a table drive test for strings.Index
// https://golang.org/pkg/strings/#Index
// Use the following test conditions
// |------------------------------------------------|
// | Value                     | Substring | Answer |
// |===========================|===========|========|
// | "Gophers are amazing!"    | "are"     | 8      |
// | "Testing in Go is fun."   | "fun"     | 17     |
// | "The answer is 42."       | "is"      | 11     |
// |------------------------------------------------|

// Rewrite the above test for strings.Index using subtests

// Here is a simple test that tests `strings.HasSuffix`.i
// https://golang.org/pkg/strings/#HasSuffix
func Test_HasSuffix(t *testing.T) {
	value := "main.go"
	if !strings.HasSuffix(value, ".go") {
		t.Fatalf("expected %s to have suffix %s", value, ".go")
	}
}

func Test_HasPrefix(t *testing.T) {
	value := "main.go"
	prefix := "main"
	if !strings.HasPrefix(value, prefix) {
		t.Fatalf("expected %s to have prefix %s", value, prefix)
	}
}

func Test_Index(t *testing.T) {
	tt := []struct {
		Value         string
		Substring     string
		ExpectedIndex int
	}{
		{Value: "Gophers are amazing!", Substring: "are", ExpectedIndex: 8},
		{Value: "Testing in Go is fun.", Substring: "fun", ExpectedIndex: 17},
		{Value: "The answer is 42.", Substring: "is", ExpectedIndex: 11},
	}

	for _, test := range tt {
		t.Run(fmt.Sprintf("Test for index of '%s' in '%s'", test.Substring, test.Value), func(st *testing.T) {
			if index := strings.Index(test.Value, test.Substring); index != test.ExpectedIndex {
				t.Errorf("Expected index %d, got index %d", test.ExpectedIndex, index)
			}
		})
	}
}
