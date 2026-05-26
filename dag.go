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

package gse

import (
	"regexp"
)

const (
	// RatioWord ratio words and letters
	RatioWord float32 = 1.5
	// RatioWordFull full ratio words and letters
	RatioWordFull float32 = 1
)

var reEng = regexp.MustCompile(`[[:alnum:]]`)

type route struct {
	freq  float64
	index int
}

// Find find word in dictionary return word's freq, pos and existence
func (seg *Segmenter) Find(str string) (float64, string, bool) {
	_ = "STUB: not implemented"
	return 0, "", false

	// FindTFIDF find word in dictionary return word's freq, inverseFreq and existence
}

func (seg *Segmenter) FindTFIDF(str string) (float64, float64, bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

// Value find word in dictionary return word's value
func (seg *Segmenter) Value(str string) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// FindAllOccs find the all search byte start in data
func FindAllOccs(data []byte, searches []string) map[string][]int {
	_ = "STUB: not implemented"
	return nil
}

// Analyze analyze the token segment info
func (seg *Segmenter) Analyze(text []string, t1 string, by ...bool) (az []AnalyzeToken) {
	_ = "STUB: not implemented"
	return nil
}

// getDag get a directed acyclic graph (DAG) from slice of runes(containing Unicode characters)
func (seg *Segmenter) getDag(runes []rune) map[int][]int { _ = "STUB: not implemented"; return nil }

func (seg *Segmenter) calc(runes []rune) map[int]route { _ = "STUB: not implemented"; return nil }

func (seg *Segmenter) hmm(bufString string, buf []rune, reg ...*regexp.Regexp) (result []string) {
	_ = "STUB: not implemented"
	return nil
}

func (seg *Segmenter) cutDAG(str string, reg ...*regexp.Regexp) []string {
	_ = "STUB: not implemented"
	return nil
}

func (seg *Segmenter) cutDAGNoHMM(str string) []string { _ = "STUB: not implemented"; return nil }

// buf = make([]rune, 0)

func (seg *Segmenter) cutAll(str string) []string { _ = "STUB: not implemented"; return nil }

func (seg *Segmenter) cutForSearch(str string, hmm ...bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// SuggestFreq suggest the words frequency
// return a suggested frequency of a word cutted to short words.
func (seg *Segmenter) SuggestFreq(words ...string) float64 { _ = "STUB: not implemented"; return 0 }
