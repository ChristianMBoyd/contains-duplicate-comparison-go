package containsDuplicate

import (
	"flag"
	"fmt"
	"math/rand"
	"testing"
)

var (
	numsLength                     = flag.Int("numsLength", 1000, "numsLength is the length of the randomized test case used for benchmarking.  This value must be >=2")
	numsLengthTestWithoutDuplicate containsDuplicateTest
	numsLengthTestWithDuplicate    containsDuplicateTest
)

type containsDuplicateTest struct {
	nums              []int
	name              string
	containsDuplicate bool
}

func generateSingleElementTest() containsDuplicateTest {
	return containsDuplicateTest{
		nums:              []int{0},
		name:              "single element",
		containsDuplicate: false,
	}
}

func generateTwoElementTestWithDuplicate() containsDuplicateTest {
	return containsDuplicateTest{
		nums:              []int{0, 0},
		name:              "two elements with duplicate",
		containsDuplicate: true,
	}
}

func generateTwoElementTestWithoutDuplicate() containsDuplicateTest {
	return containsDuplicateTest{
		nums:              []int{0, 1},
		name:              "two elements without duplicate",
		containsDuplicate: false,
	}
}

// generateNELementTestWithoutDuplicate creates a test case without
// duplicates, where the length of the input nums is determined by
// the command line flag numsLength.  The values used are randomly
// generated and restricted by the maximal values of Go's int.
func generateNElementTestWithoutDuplicate() containsDuplicateTest {
	var nums []int
	previousValues := make(map[int]bool)
	for i := range *numsLength {
		value := rand.Int()
		if value%2 == 0 {
			value *= -1 // convert even values to negatives
		}

		if previousValues[value] {
			fmt.Printf("Repeated value (%d) at index %d!\n", value, i)
		}
		previousValues[value] = true

		nums = append(nums, value)
	}

	return containsDuplicateTest{
		nums:              nums,
		name:              fmt.Sprintf("%d elements without duplicate", *numsLength),
		containsDuplicate: false,
	}
}

// generateNElementTestWithDuplicate creates a test case with at
// least one duplicate, where the length of the input nums is
// determined by the command line flag numsLength.  The values used
// are randomly generated and restricted by the maximal values of
// Go's int.
func generateNElementTestWithDuplicate() containsDuplicateTest {
	testWithoutDuplicate := generateNElementTestWithoutDuplicate()
	nums := make([]int, *numsLength)
	copy(nums, testWithoutDuplicate.nums)

	i, j := generateUniqueNonnegativePair(*numsLength)
	if i == j {
		fmt.Printf("Generated identical index pair (%d, %d) when trying to make a duplicate!", i, j)
	}
	nums[j] = nums[i]
	return containsDuplicateTest{
		nums:              nums,
		name:              fmt.Sprintf("%d elements with at least one duplicate", *numsLength),
		containsDuplicate: true,
	}
}

// generateUniqueNonnegativePair creates a nonnegative pair of
// integers within the interval [0, max).  If max is less than or
// equal to zero, then this function will fail.
func generateUniqueNonnegativePair(max int) (int, int) {
	i := rand.Int() % max
	j := rand.Int() % max

	// Because of the modulus, a small max has a high probability of
	// duplicate indices despite the "random" selection
	if i == j {
		j = (max - i) % max // attempt to move j away from i
	}
	if i == j {
		j = (i + 1) % max // desperate attempt to move j off i
	}

	return i, j
}

// generateContainsDuplicateTests collects the tests cases used to
// test the validity of the various solutions to the "Contains
// Duplicate" prompt.  In particular, copies of the numsLength tests
// are created so that solutions are allowed to mutate the test data.
func generateContainsDuplicateTests() []containsDuplicateTest {
	numsWithDuplicate := copyNumsFromTest(&numsLengthTestWithDuplicate)
	numsWithoutDuplicate := copyNumsFromTest(&numsLengthTestWithoutDuplicate)

	return []containsDuplicateTest{
		generateSingleElementTest(),
		generateTwoElementTestWithDuplicate(),
		generateTwoElementTestWithoutDuplicate(),
		{
			nums:              numsWithDuplicate,
			name:              numsLengthTestWithDuplicate.name,
			containsDuplicate: numsLengthTestWithDuplicate.containsDuplicate,
		},
		{
			nums:              numsWithoutDuplicate,
			name:              numsLengthTestWithoutDuplicate.name,
			containsDuplicate: numsLengthTestWithoutDuplicate.containsDuplicate,
		},
	}
}

func copyNumsFromTest(test *containsDuplicateTest) []int {
	nums := make([]int, len(test.nums))
	copy(nums, test.nums)
	return nums
}

func TestThatCreatesTheDynamicSizedTestCases(t *testing.T) {
	// The command line flags are only passed just before the tests
	// run.  If these were assigned in the var() section, then they
	// would always use the default value of numsLength.
	numsLengthTestWithDuplicate = generateNElementTestWithDuplicate()
	numsLengthTestWithoutDuplicate = generateNElementTestWithoutDuplicate()
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
	numsWithDuplicate := copyNumsFromTest(&numsLengthTestWithDuplicate)
	numsWithoutDuplicate := copyNumsFromTest(&numsLengthTestWithoutDuplicate)
	for range b.N {
		UsingMap(numsWithDuplicate)
		UsingMap(numsWithoutDuplicate)
	}
}

func BenchmarkUsingLeetCodeSolution(b *testing.B) {
	numsWithDuplicate := copyNumsFromTest(&numsLengthTestWithDuplicate)
	numsWithoutDuplicate := copyNumsFromTest(&numsLengthTestWithoutDuplicate)
	for range b.N {
		LeetCodeSolution(numsWithDuplicate)
		LeetCodeSolution(numsWithoutDuplicate)
	}
}
