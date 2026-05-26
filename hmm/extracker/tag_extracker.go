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
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/relevance"
	"github.com/go-ego/gse/hmm/segment"
	"github.com/go-ego/gse/types"
)

// TagExtracter is extract tags struct.
type TagExtracter struct {
	seg gse.Segmenter

	// calculate weight by Relevance(including IDF,TF-IDF,BM25 and so on)
	Relevance relevance.Relevance
	// stopWord *stopwords.StopWord
}

// WithGse register the gse segmenter
func (t *TagExtracter) WithGse(segs gse.Segmenter) {
	_ = "STUB: not implemented"

	// LoadDict load and create a new dictionary from the file
	return
}

func (t *TagExtracter) LoadDict(fileName ...string) error { _ = "STUB: not implemented"; return nil }

// LoadIdf load and create a new Idf dictionary from the file.
func (t *TagExtracter) LoadIdf(fileName ...string) error { _ = "STUB: not implemented"; return nil }

// LoadIdfStr load and create a new Idf dictionary from the string.
func (t *TagExtracter) LoadIdfStr(str string) error { _ = "STUB: not implemented"; return nil }

// LoadTFIDF load and create a new TFIDF dictionary from the file.
func (t *TagExtracter) LoadTFIDF(fileName ...string) error { _ = "STUB: not implemented"; return nil }

// LoadBM25 load and create a new BM25 dictionary from the file.
// params setting: the k1 and b to defind for calcluate bm25
//
//	types.BM25Setting{
//		K1
//	 	B
//	}
//
// params: fileList:
//
//	type LoadBM25DictFile struct {
//			FilePath string
//			FileType int
//			Number   float64
//	}
func (t *TagExtracter) LoadBM25(setting *types.BM25Setting, fileList []*types.LoadDictFile) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// load dict file and corpus file

// Distinguishing dictionary types

// LoadStopWords load and create a new StopWord dictionary from the file.
func (t *TagExtracter) LoadStopWords(fileName ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// ExtractTags extract the topK keywords from text.
func (t *TagExtracter) ExtractTags(text string, topK int) (tags segment.Segments) {
	_ = "STUB: not implemented"
	return *

	// If no correlation algorithm, we will set the idf for default.
	new(segment.Segments)
}

// handler text to construct segment with weight

// sort by weight desc

// choose the top keywords if length of weightSeg bigger than topK
