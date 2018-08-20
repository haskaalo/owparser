package owparser

import (
	"log"
	"os"
	"testing"
)

var resultAll *All

func BenchmarkNewAll(b *testing.B) {
	dir, _ := os.Getwd()
	testfile, err := os.Open(dir + "/testfiles/trev.html")
	if err != nil {
		log.Panicln(err)
	}
	careerProfile, err := NewCareerProfile(testfile)
	if err != nil {
		log.Panicln(err)
	}
	var result *All

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		result = careerProfile.NewAll()
	}
	resultAll = result
}
