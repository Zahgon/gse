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
Package hmm is the Golang HMM cut module
*/
package hmm

import (
	"regexp"
)

var (
	regHan  = regexp.MustCompile(`\p{Han}+`)
	regSkip = regexp.MustCompile(`(\d+\.\d+|[a-zA-Z0-9]+)`)
)

// func LoadFile(filePath string) map[rune]float64 {
//
// }

// LoadModel load the HMM model
func LoadModel(prob ...map[rune]float64) { _ = "STUB: not implemented"; return }

func internalCut(text string) []string { _ = "STUB: not implemented"; return nil }

// Cut cuts text to words using HMM with Viterbi algorithm
func Cut(text string, reg ...*regexp.Regexp) []string { _ = "STUB: not implemented"; return nil }

// find(text, cuts, cutLoc, nonCutLoc)

func locJudge(str string, cutLoc, nonCutLoc []int) (loc []int) {
	_ = "STUB: not implemented"
	return nil
}
