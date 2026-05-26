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

package extracker

import (
	"sort"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
	"github.com/go-ego/gse/hmm/segment"
)

const dampingFactor = 0.85

var (
	defaultAllowPOS = []string{"ns", "n", "vn", "v"}
)

// TextRanker is extract tags struct.
type TextRanker struct {
	seg pos.Segmenter
	HMM bool
}

// WithGse register the gse segmenter
func (t *TextRanker) WithGse(segs gse.Segmenter) { _ = "STUB: not implemented"; return }

// LoadDict load and create a new dictionary from the file for Textranker
func (t *TextRanker) LoadDict(fileName ...string) error {
	_ = "STUB: not implemented"
	// t.seg = new(pos.Segmenter)
	return nil
}

type edge struct {
	start, end string
	weight     float64
}

type edges []edge

func (es edges) Len() int { _ = "STUB: not implemented"; return 0 }

func (es edges) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (es edges) Swap(i, j int) { _ = "STUB: not implemented"; return }

type undirectWeightedGraph struct {
	graph map[string]edges
	keys  sort.StringSlice
}

func newUndirectWeightedGraph() *undirectWeightedGraph { _ = "STUB: not implemented"; return nil }

func (u *undirectWeightedGraph) addEdge(start, end string, weight float64) {
	_ = "STUB: not implemented"
	// # use a tuple (start, end, weight) instead of a Edge object
	return
}

func (u *undirectWeightedGraph) rank() segment.Segments {
	_ = "STUB: not implemented"
	return *new(segment.Segments)
}

// TextRankWithPOS extracts keywords from text using TextRank algorithm.
// Parameter allowPOS allows a []string pos list.
func (t *TextRanker) TextRankWithPOS(text string, topK int, allowPOS []string) segment.Segments {
	_ = "STUB: not implemented"
	return *new(segment.Segments)
}

// TextRank extract keywords from text using TextRank algorithm.
// Parameter topK specify how many top keywords to be returned at most.
func (t *TextRanker) TextRank(text string, topK int) segment.Segments {
	_ = "STUB: not implemented"
	return *new(segment.Segments)
}
