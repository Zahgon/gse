// Copyright 2013 Hui Chen
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

// ToString converts a segments slice to string return the string
//
//	 two output modes:
//
//		normal mode (searchMode=false）
//		search mode（searchMode=true）
//
// default searchMode=false
// search mode is used search engine, and will output more results
func ToString(segs []Segment, searchMode ...bool) (output string) {
	_ = "STUB: not implemented"
	return ""
}

func tokenToString(token *Token) (output string) { _ = "STUB: not implemented"; return "" }

func tokenToBytes(token *Token) (output []byte) { _ = "STUB: not implemented"; return nil }

// ToSlice converts a segments to slice return string slice
func ToSlice(segs []Segment, searchMode ...bool) (output []string) {
	_ = "STUB: not implemented"
	return nil
}

func tokenToSlice(token *Token) (output []string) { _ = "STUB: not implemented"; return nil }

// ToPos converts a segments slice to []SegPos
func ToPos(segs []Segment, searchMode ...bool) (output []SegPos) {
	_ = "STUB: not implemented"
	return nil
}

func tokenToPos(token *Token) (output []SegPos) { _ = "STUB: not implemented"; return nil }

// let make multiple []Text into one string output
func textToString(text []Text) (output string) { _ = "STUB: not implemented"; return "" }

// let make []Text toString returns a string output
func textSliceToString(text []Text) string {
	_ = "STUB: not implemented"

	// return total length of text slice
	return ""
}

func textSliceByteLen(text []Text) (length int) { _ = "STUB: not implemented"; return 0 }

func textSliceToBytes(text []Text) []byte { _ = "STUB: not implemented"; return nil }

// Join is better string splicing
func Join(text []Text) string { _ = "STUB: not implemented"; return "" }

// Special case for common small values.
// Remove if github.com/golang/go/issues/6714 is fixed

// Special case for common small values.
// Remove if #6714 is fixed

func printTokens(tokens []*Token, numTokens int) (output string) {
	_ = "STUB: not implemented"
	return ""
}

func toWords(strings ...string) []Text { _ = "STUB: not implemented"; return nil }

func bytesToString(bytes []Text) (output string) { _ = "STUB: not implemented"; return "" }
