package stopwords

import (
	"github.com/go-ego/gse"
)

// StopWordMap the default stop words.
var StopWordMap = map[string]bool{
	"the":   true,
	"of":    true,
	"is":    true,
	"and":   true,
	"to":    true,
	"in":    true,
	"that":  true,
	"we":    true,
	"for":   true,
	"an":    true,
	"are":   true,
	"by":    true,
	"be":    true,
	"as":    true,
	"on":    true,
	"with":  true,
	"can":   true,
	"if":    true,
	"from":  true,
	"which": true,
	"you":   true,
	"it":    true,
	"this":  true,
	"then":  true,
	"at":    true,
	"have":  true,
	"all":   true,
	"not":   true,
	"one":   true,
	"has":   true,
	"or":    true,
}

// StopWord is a dictionary for all stop words.
type StopWord struct {
	stopWordMap map[string]bool

	seg gse.Segmenter
}

// AddStop add a token to StopWord dictionary.
func (s *StopWord) AddStop(text string) { _ = "STUB: not implemented"; return }

// RemoveStop remove a token from StopWord dictionary.
func (s *StopWord) RemoveStop(text string) { _ = "STUB: not implemented"; return }

// NewStopWord create a new StopWord with the default stop words.
func NewStopWord() *StopWord { _ = "STUB: not implemented"; return nil }

// IsStopWord check the word is a stop word
func (s *StopWord) IsStopWord(word string) bool { _ = "STUB: not implemented"; return false }

// LoadDict load the idf stop dictionary
func (s *StopWord) LoadDict(files ...string) error { _ = "STUB: not implemented"; return nil }
