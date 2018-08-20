package owparser

import (
	"log"
	"os"
	"testing"
)

var resultCareerProfile *CareerProfile

func BenchmarkNewCareerProfile(b *testing.B) {
	dir, _ := os.Getwd()
	testfile, err := os.Open(dir + "/testfiles/trev.html")
	if err != nil {
		log.Panicln(err)
	}
	b.ResetTimer()
	var result *CareerProfile
	for n := 0; n < b.N; n++ {
		// record result to prevent the compiler eliminating the function call.
		result, _ = NewCareerProfile(testfile)
	}
	resultCareerProfile = result
}
