// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package relevance

import (
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/segment"
)

// TFIDF a measure of importance of a word to a document in a collection.
// Term Frequency-Inverse Document Frequency
// ref: https://en.wikipedia.org/wiki/Tf–idf
type TFIDF struct {
	// the list of word frequencies
	freqs []float64

	Base
}

// AddToken add a new word with TFIDF into the dictionary.
func (t *TFIDF) AddToken(text string, freq float64, pos ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadStopWord load stop word for TFIDF
func (t *TFIDF) LoadStopWord(fileName ...string) error { _ = "STUB: not implemented"; return nil }

// LoadDict load dict for TFIDF seg
func (t *TFIDF) LoadDict(files ...string) error { _ = "STUB: not implemented"; return nil }

// LoadDictStr load dict for TFIDF seg
func (t *TFIDF) LoadDictStr(dictStr string) error { _ = "STUB: not implemented"; return nil }

// Freq return the TFIDF of the word
func (t *TFIDF) Freq(key string) (float64, interface{}, bool) {
	_ = "STUB: not implemented"
	return 0, nil,

		// NumTokens return the TFIDF tokens' num
		false
}

func (t *TFIDF) NumTokens() int { _ = "STUB: not implemented"; return 0 }

// TotalFreq return the TFIDF total frequency
func (t *TFIDF) TotalFreq() float64 { _ = "STUB: not implemented"; return 0 }

// FreqMap return the TFIDF freq map
func (t *TFIDF) FreqMap(text string) map[string]float64 { _ = "STUB: not implemented"; return nil }

// calculateIdf calculate the word's weight by TFIDF
func (t *TFIDF) calculateWeight(term string) float64 { _ = "STUB: not implemented"; return 0 }

// ConstructSeg construct segment with weight
func (t *TFIDF) ConstructSeg(text string) segment.Segments {
	_ = "STUB: not implemented"
	// make segment list by total freq num
	return *new(segment.Segments)
}

// GetSeg get TFIDF Segmenter
func (t *TFIDF) GetSeg() gse.Segmenter {
	_ = "STUB: not implemented"

	// LoadCorpus tf idf no need to load corpus
	return *new(gse.Segmenter)
}

func (t *TFIDF) LoadCorpus(path ...string) error {
	_ = "STUB: not implemented"

	// NewTFIDF create a new TFIDF
	return nil
}

func NewTFIDF() Relevance { _ = "STUB: not implemented"; return *new(Relevance) }
