package owparser

import (
	"log"
	"os"
	"testing"
)

var resultStats *Stats

func BenchmarkNewStats(b *testing.B) {
	dir, _ := os.Getwd()
	testfile, err := os.Open(dir + "/testfiles/trev.html")
	if err != nil {
		log.Panicln(err)
	}
	careerProfile, err := NewCareerProfile(testfile)
	if err != nil {
		log.Panicln(err)
	}
	var result *Stats
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		// record result to prevent the compiler eliminating the function call.
		result = careerProfile.NewStats(QuickPlay)
	}
	resultStats = result
}
