package dog

import (
	"fmt"
	"testing"
)

//test
func TestYears(t *testing.T) {
	test := Years(10)
	if test != 70 {
		t.Error("Got", test, "wanted:", 70)
	}
}

//Example for docs.
func ExampleYears() {
	fmt.Println(Years(10))
	// Output: 70
}

//benchmark
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

//test
func TestYearsTwo(t *testing.T) {
	test := YearsTwo(10)
	if test != 70 {
		t.Error("Got", test, "wanted:", 70)
	}
}

//Example for docs.
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output: 70
}

//benchmark
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
