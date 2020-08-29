package dog

import (
	"testing"
)

func BenchmarkTodogYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//fmt.Println(ToDogYears(10)) cool one to try and see different output
		ToDogYears(10)
	}
}
