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

/*
Package gse Go efficient multilingual NLP and text segmentation
*/
package gse

import (
	"regexp"
)

const (
	// Version get the gse version
	Version = "v1.0.1.705, Green Lake!"

	// minTokenFrequency = 2 // only read tokens with frequency >= 2 from the dictionary
)

// GetVersion get the version of gse
func GetVersion() string {
	_ = "STUB: not implemented"

	// Prob define the hmm model struct
	return ""
}

type Prob struct {
	B, E, M, S map[rune]float64
}

// New return a new gse segmenter
func New(files ...string) (seg Segmenter, err error) {
	_ = "STUB: not implemented"
	return *new(Segmenter), nil
}

// Cut cuts a str into words using accurate mode.
// Parameter hmm controls whether to use the HMM(Hidden Markov Model)
// or use the user's model.
//
// seg.Cut(text):
//
//	use the shortest path
//
// seg.Cut(text, false):
//
//	use cut dag not hmm
//
// seg.Cut(text, true):
//
//	use cut dag and hmm mode
func (seg *Segmenter) Cut(str string, hmm ...bool) []string { _ = "STUB: not implemented"; return nil }

// return seg.cutDAGNoHMM(str)

// CutSearch cuts str into words using search engine mode.
func (seg *Segmenter) CutSearch(str string, hmm ...bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// CutAll cuts a str into words using full mode.
func (seg *Segmenter) CutAll(str string) []string { _ = "STUB: not implemented"; return nil }

// CutDAG cut string with DAG use hmm and regexp
func (seg *Segmenter) CutDAG(str string, reg ...*regexp.Regexp) []string {
	_ = "STUB: not implemented"
	return nil
}

// CutDAGNoHMM cut string with DAG not use hmm
func (seg *Segmenter) CutDAGNoHMM(str string) []string { _ = "STUB: not implemented"; return nil }

// CutStr cut []string with Cut return string
func (seg *Segmenter) CutStr(str []string, separator ...string) (r string) {
	_ = "STUB: not implemented"
	return ""
}

// LoadModel load the hmm model (default is Chinese char)
//
// Use the user's model:
//
//	seg.LoadModel(B, E, M, S map[rune]float64)
func (seg *Segmenter) LoadModel(prob ...map[rune]float64) { _ = "STUB: not implemented"; return }

// HMMCut cut sentence string use HMM with Viterbi
func (seg *Segmenter) HMMCut(str string, reg ...*regexp.Regexp) []string {
	_ = "STUB: not implemented"
	// hmm.LoadModel(prob...)
	return nil
}

// HMMCutMod cut sentence string use HMM with Viterbi
func (seg *Segmenter) HMMCutMod(str string, prob ...map[rune]float64) []string {
	_ = "STUB: not implemented"
	return nil
}

// Slice use modeSegment segment return []string
// using search mode if searchMode is true
func (seg *Segmenter) Slice(s string, searchMode ...bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// Slice use modeSegment segment return string
// using search mode if searchMode is true
func (seg *Segmenter) String(s string, searchMode ...bool) string {
	_ = "STUB: not implemented"
	return ""
}

// SegPos type a POS struct
type SegPos struct {
	Text, Pos string
}

// Pos return text and pos array
func (seg *Segmenter) Pos(s string, searchMode ...bool) []SegPos {
	_ = "STUB: not implemented"
	return nil
}

// PosStr cut []SegPos with Pos return string
func (seg *Segmenter) PosStr(str []SegPos, separator ...string) (r string) {
	_ = "STUB: not implemented"
	return ""
}
