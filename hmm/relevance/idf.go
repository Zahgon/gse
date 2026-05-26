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

// Idf type a dictionary for all words with the
// IDFs(Inverse Document Frequency).
type Idf struct {
	// median of word frequencies for calculate the weight of backup
	median float64

	// the list of word frequencies
	freqs []float64

	Base
}

// AddToken add a new word with IDF into the dictionary.
func (i *Idf) AddToken(text string, freq float64, pos ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadDict load the idf dictionary
func (i *Idf) LoadDict(files ...string) error { _ = "STUB: not implemented"; return nil }

// LoadStopWord load stop word for IDF
func (i *Idf) LoadStopWord(fileName ...string) error { _ = "STUB: not implemented"; return nil }

// LoadDictStr load dict for IDF seg
func (i *Idf) LoadDictStr(dictStr string) error { _ = "STUB: not implemented"; return nil }

// Freq return the IDF of the word
func (i *Idf) Freq(key string) (float64, interface{}, bool) {
	_ = "STUB: not implemented"
	return 0,

		// NumTokens return the IDF tokens' num
		nil, false
}

func (i *Idf) NumTokens() int { _ = "STUB: not implemented"; return 0 }

// TotalFreq return the IDF total frequency
func (i *Idf) TotalFreq() float64 { _ = "STUB: not implemented"; return 0 }

// FreqMap return the IDF freq map
func (i *Idf) FreqMap(text string) map[string]float64 { _ = "STUB: not implemented"; return nil }

// calculateWeight calculate the word's weight by IDF
func (i *Idf) calculateWeight(k string, v float64) float64 { _ = "STUB: not implemented"; return 0 }

// ConstructSeg construct segment with weight
func (i *Idf) ConstructSeg(text string) segment.Segments {
	_ = "STUB: not implemented"
	// make segment list by total freq num
	return *new(segment.Segments)
}

// GetSeg get IDF Segmenter
func (i *Idf) GetSeg() gse.Segmenter {
	_ = "STUB: not implemented"

	// LoadCorpus idf no need to load corpus
	return *new(gse.Segmenter)
}

func (i *Idf) LoadCorpus(path ...string) error {
	_ = "STUB: not implemented"

	// NewIdf create a new Idf
	return nil
}

func NewIdf() Relevance { _ = "STUB: not implemented"; return *new(Relevance) }
