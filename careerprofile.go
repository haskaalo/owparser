package owparser

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

// CareerProfile ...
type CareerProfile struct {
	document *goquery.Document
}

// NewCareerProfile Get a CareerProfile struct from a Reader
func NewCareerProfile(reader io.Reader) (*CareerProfile, error) {
	career := new(CareerProfile)

	var err error
	career.document, err = goquery.NewDocumentFromReader(reader)

	return career, err
}
