// Copyright 2016 The go-ego Project Developers.
//
// See the COPYRIGHT file at the top-level directory of this distribution and at
// https://github.com/go-ego/gse/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0>
//
// This file may not be copied, modified, or distributed
// except according to those terms.

package gse

import (
	"github.com/go-ego/gse/types"
)

// //go:embed data/dict/dictionary.txt
// var dataDict string

// NewEmbed return new gse segmenter by embed dictionary
func NewEmbed(dict ...string) (seg Segmenter, err error) {
	_ = "STUB: not implemented"
	return *new(Segmenter), nil
}

func (seg *Segmenter) loadZh() error { _ = "STUB: not implemented"; return nil }

func (seg *Segmenter) loadZhST(d string) (begin int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// err = seg.LoadDictStr(dataDict)

// LoadDictEmbed load the dictionary by embed file
func (seg *Segmenter) LoadDictEmbed(dict ...string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// return seg.LoadDictStr(dataDict)

// LoadDictStr load the dictionary from dict path
func (seg *Segmenter) LoadDictStr(dict string) error { _ = "STUB: not implemented"; return nil }

// add the words to the token

// LoadTFIDFDictStr load the TFIDF dictionary from dict path
func (seg *Segmenter) LoadTFIDFDictStr(dictFile *types.LoadDictFile) error {
	_ = "STUB: not implemented"
	return nil
}

// frequency

// invserse frequency

// add the words to the token

// LoadStopEmbed load the stop dictionary from embed file
func (seg *Segmenter) LoadStopEmbed(dict ...string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// LoadStopStr load the stop dictionary from dict path
func (seg *Segmenter) LoadStopStr(dict string) error { _ = "STUB: not implemented"; return nil }
