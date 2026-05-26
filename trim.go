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

func notPunct(ru []rune) bool { _ = "STUB: not implemented"; return false }

// TrimPunct trim []string exclude space and punct
func (seg *Segmenter) TrimPunct(s []string) (r []string) { _ = "STUB: not implemented"; return nil }

// TrimPosPunct trim SegPos not space and punct
func (seg *Segmenter) TrimPosPunct(se []SegPos) (re []SegPos) {
	_ = "STUB: not implemented"
	return nil
}

// TrimWithPos trim some seg with pos
func (seg *Segmenter) TrimWithPos(se []SegPos, pos ...string) (re []SegPos) {
	_ = "STUB: not implemented"
	return nil
}

// Stop trim []string stop word
func (seg *Segmenter) Stop(s []string) (r []string) { _ = "STUB: not implemented"; return nil }

// Trim trim []string exclude symbol, space and punct
func (seg *Segmenter) Trim(s []string) (r []string) { _ = "STUB: not implemented"; return nil }

// TrimSymbol trim []string exclude symbol, space and punct
func (seg *Segmenter) TrimSymbol(s []string) (r []string) { _ = "STUB: not implemented"; return nil }

// TrimPos trim SegPos not symbol, space and punct
func (seg *Segmenter) TrimPos(s []SegPos) (r []SegPos) { _ = "STUB: not implemented"; return nil }

// CutStop cut string and tirm stop
func (seg *Segmenter) CutStop(str string, hmm ...bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// CutTrim cut string and tirm
func (seg *Segmenter) CutTrim(str string, hmm ...bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// PosTrim cut string pos and trim
func (seg *Segmenter) PosTrim(str string, search bool, pos ...string) []SegPos {
	_ = "STUB: not implemented"
	return nil
}

// PosTrimArr cut string return pos.Text []string
func (seg *Segmenter) PosTrimArr(str string, search bool, pos ...string) (re []string) {
	_ = "STUB: not implemented"
	return nil
}

// PosTrimStr cut string return pos.Text string
func (seg *Segmenter) PosTrimStr(str string, search bool, pos ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// CutTrimHtml cut string trim html and symbol return []string
func (seg *Segmenter) CutTrimHtml(str string, hmm ...bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// CutTrimHtmls cut string trim html and symbol return string
func (seg *Segmenter) CutTrimHtmls(str string, hmm ...bool) string {
	_ = "STUB: not implemented"
	return ""
}

// CutUrl cut url string trim symbol return []string
func (seg *Segmenter) CutUrl(str string, num ...bool) []string {
	_ = "STUB: not implemented"

	// seg.Num = true
	return nil
}

// CutUrls cut url string trim symbol return string
func (seg *Segmenter) CutUrls(str string, num ...bool) string { _ = "STUB: not implemented"; return "" }

// SplitNum cut string by num to []string
func SplitNum(text string) []string { _ = "STUB: not implemented"; return nil }

// SplitNums cut string by num to string
func SplitNums(text string) string { _ = "STUB: not implemented"; return "" }

// FilterEmoji filter the emoji
func FilterEmoji(text string) (new string) { _ = "STUB: not implemented"; return "" }

// FilterSymbol filter the symbol
func FilterSymbol(text string) (new string) { _ = "STUB: not implemented"; return "" }

// FilterHtml filter the html tag
func FilterHtml(text string) string { _ = "STUB: not implemented"; return "" }

// FilterLang filter the language
func FilterLang(text, lang string) (new string) { _ = "STUB: not implemented"; return "" }

// Range range text to []string
func Range(text string) (new []string) { _ = "STUB: not implemented"; return nil }

// RangeText range text to string
func RangeText(text string) (new string) { _ = "STUB: not implemented"; return "" }
