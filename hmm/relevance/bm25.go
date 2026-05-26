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
	"github.com/go-ego/gse/types"
)

// BM25 Best Match
// ref: https://en.wikipedia.org/wiki/Okapi_BM25
type BM25 struct {
	// K1 Saturation Parameter
	// Controls the saturation of the TF,
	// i.e. a word frequency that exceeds the value of this parameter is not given more weight.
	// A lower k1 will make the word frequency less influential
	// and a higher k1 will make the word frequency more influential.
	// if not defined K1 , we will define it in 1.25
	K1 float64

	// B Length Normalization Parameter
	// Controls the degree of normalization of document length.
	// A lower b will make shorter documents more important
	// and a higher b will make longer documents more important.
	// so and K1 , if not defined by client, we will define it in 0.75
	B float64

	// AverageDocLength Average Document Length
	// Indicates the average vocabulary per document in the entire document set.
	// This value is used to normalize the document length in order to compare documents of different lengths.
	AverageDocLength float64

	// Base default setting
	Base
}

// AddToken add a new word with TFIDF into the dictionary.
func (bm25 *BM25) AddToken(text string, freq float64, pos ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadStopWord load stop word for TFIDF
func (bm25 *BM25) LoadStopWord(fileName ...string) error { _ = "STUB: not implemented"; return nil }

// LoadDict load dict for TFIDF seg
func (bm25 *BM25) LoadDict(files ...string) error { _ = "STUB: not implemented"; return nil }

// bm25 needs tf and idf value , so we just get the tfidf path and loading it.

// LoadDictStr load dict for BM25 seg
func (bm25 *BM25) LoadDictStr(dictStr string) error { _ = "STUB: not implemented"; return nil }

// Freq return the BM25 of the word
// BM25 need TF and IDF value, so we just use FindTFIDF func
func (bm25 *BM25) Freq(key string) (float64, interface{}, bool) {
	_ = "STUB: not implemented"
	return 0, nil, false

	// NumTokens return the BM25 tokens' num
}

func (bm25 *BM25) NumTokens() int { _ = "STUB: not implemented"; return 0 }

// TotalFreq return the BM25 total frequency
func (bm25 *BM25) TotalFreq() float64 { _ = "STUB: not implemented"; return 0 }

// FreqMap return the BM25 freq map
func (bm25 *BM25) FreqMap(text string) map[string]float64 { _ = "STUB: not implemented"; return nil }

// calculateK Calculate the K value for bm25
func (bm25 *BM25) calculateK(docNum float64) float64 { _ = "STUB: not implemented"; return 0 }

// calculateWeight calculate the word's weight by BM25
func (bm25 *BM25) calculateWeight(term string) float64 { _ = "STUB: not implemented"; return 0 }

// ConstructSeg construct segment with weight
func (bm25 *BM25) ConstructSeg(text string) segment.Segments {
	_ = "STUB: not implemented"
	// make segment list by total freq num
	return *new(segment.Segments)
}

// GetSeg get TFIDF Segmenter
func (bm25 *BM25) GetSeg() gse.Segmenter {
	_ = "STUB: not implemented"

	// LoadCorpus for calculate the average length of corpus
	return *new(gse.Segmenter)
}

func (bm25 *BM25) LoadCorpus(path ...string) (err error) { _ = "STUB: not implemented"; return nil }

// NewBM25 create a new BM25
func NewBM25(bm25Setting *types.BM25Setting) Relevance {
	_ = "STUB: not implemented"
	// init value
	return *new(Relevance)
}
