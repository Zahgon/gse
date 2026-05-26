// Copyright 2013 Hui Chen
// Copyright 2016 ego authors
//
// Copyright 2016 The go-ego Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-ego/gse/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0>
//
// This file may not be copied, modified, or distributed
// except according to those terms.

package gse

import (
	"github.com/vcaesar/cedar"
)

// Dictionary struct implements a string double array trie.
// one segment maybe in leaf node or not
type Dictionary struct {
	trie *cedar.Cedar // Cedar double array trie

	maxTokenLen int     // the maximum length of the dictionary
	Tokens      []Token // the all tokens in the dictionary, to traverse
	totalFreq   float64 // the total number of tokens in the dictionary
}

// NewDict a new dictionary trie
func NewDict() *Dictionary { _ = "STUB: not implemented"; return nil }

// MaxTokenLen the maximum length of the dictionary
func (dict *Dictionary) MaxTokenLen() int { _ = "STUB: not implemented"; return 0 }

// NumTokens the number of tokens in the dictionary
func (dict *Dictionary) NumTokens() int { _ = "STUB: not implemented"; return 0 }

// TotalFreq the total frequency of the dictionary
func (dict *Dictionary) TotalFreq() float64 { _ = "STUB: not implemented"; return 0 }

// AddToken add a token to the dictionary
func (dict *Dictionary) AddToken(token Token) error { _ = "STUB: not implemented"; return nil }

// RemoveToken remove token in dictionary
func (dict *Dictionary) RemoveToken(token Token) error { _ = "STUB: not implemented"; return nil }

// LookupTokens finds tokens and words in the dictionary, matching the given pattern
// and returns the number of tokens
func (dict *Dictionary) LookupTokens(
	words []Text, tokens []*Token) (numOfTokens int) {
	_ = "STUB: not implemented"
	return 0
}

// Find find the word in the dictionary is non-existent
// and the word's frequency and pos
func (dict *Dictionary) Find(word []byte) (float64, string, bool) {
	_ = "STUB: not implemented"
	return 0, "", false
}

func (dict *Dictionary) FindTFIDF(word []byte) (float64, float64, bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

// Value find word in the dictionary
// return the word's value and id
func (dict *Dictionary) Value(word []byte) (val, id int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
