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

package pos

import (
	"regexp"

	"github.com/go-ego/gse"
)

var (
	reHanDetail  = regexp.MustCompile(`(\p{Han}+)`)
	reSkipDetail = regexp.MustCompile(`([[\.[:digit:]]+|[:alnum:]]+)`)

	reEng  = regexp.MustCompile(`[[:alnum:]]`)
	reNum  = regexp.MustCompile(`[\.[:digit:]]+`)
	reEng1 = regexp.MustCompile(`[[:alnum:]]$`)

	reHanInternal  = regexp.MustCompile(`([\p{Han}+[:alnum:]+#&\._]+)`)
	reSkipInternal = regexp.MustCompile(`(\r\n|\s)`)
)

// SegPos type POS struct
type SegPos struct {
	Text, Pos string
}

// Segmenter is a segmentation struct
type Segmenter struct {
	dict Dict
}

// WithGse register the gse segmenter
func (seg *Segmenter) WithGse(segs gse.Segmenter) { _ = "STUB: not implemented"; return }

// LoadDict load dictionary from the file.
func (seg *Segmenter) LoadDict(fileName ...string) error { _ = "STUB: not implemented"; return nil }

func (seg *Segmenter) cutDetailInternal(text string) (result []gse.SegPos) {
	_ = "STUB: not implemented"
	return nil
}

func (seg *Segmenter) cutDetail(text string) (result []gse.SegPos) {
	_ = "STUB: not implemented"
	return nil
}

func (seg *Segmenter) getDag(runes []rune) map[int][]int { _ = "STUB: not implemented"; return nil }

type route struct {
	freq  float64
	index int
}

func (seg *Segmenter) calc(runes []rune) map[int]route { _ = "STUB: not implemented"; return nil }

type cutFunc func(text string) []gse.SegPos

func (seg *Segmenter) cutDAG(text string) (result []gse.SegPos) {
	_ = "STUB: not implemented"
	return nil
}

func (seg *Segmenter) bufn(buf []rune) (result []gse.SegPos) { _ = "STUB: not implemented"; return nil }

func (seg *Segmenter) cutDAGNoHMM(text string) (result []gse.SegPos) {
	_ = "STUB: not implemented"
	return nil
}

// buf = make([]rune, 0)

// Cut cuts a text into words.
// Parameter hmm controls whether to use the HMM
func (seg *Segmenter) Cut(text string, hmm ...bool) (result []gse.SegPos) {
	_ = "STUB: not implemented"
	return nil
}

// TrimPunct not space and punct
func (seg *Segmenter) TrimPunct(se []gse.SegPos) (re []gse.SegPos) {
	_ = "STUB: not implemented"
	return nil
}

// Trim not space and punct
func (seg *Segmenter) Trim(se []gse.SegPos) (re []gse.SegPos) {
	_ = "STUB: not implemented"
	return nil
}

// TrimWithPos trim some pos
func (seg *Segmenter) TrimWithPos(se []gse.SegPos, pos ...string) (re []gse.SegPos) {
	_ = "STUB: not implemented"
	return nil
}
