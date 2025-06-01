package containsDuplicate

import (
	"testing"
)

type containsDuplicateTest struct {
	nums              []int
	name              string
	containsDuplicate bool
}

func generateContainsDuplicateTests() []containsDuplicateTest {
	return []containsDuplicateTest{
		{
			nums:              []int{0},
			name:              "single element",
			containsDuplicate: false,
		},
	}
}

func TestUsingMap(t *testing.T) {
	for _, test := range generateContainsDuplicateTests() {
		t.Run(test.name, func(t *testing.T) {
			if containsDuplicate := UsingMap(test.nums); containsDuplicate != test.containsDuplicate {
				t.Errorf("UsingMap returned %v when the correct answer was %v", containsDuplicate, test.containsDuplicate)
			}
		})
	}
}

func BenchmarkUsingMap(b *testing.B) {
	// Check the performance
}
