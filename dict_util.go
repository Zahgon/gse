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
	"bufio"

	"github.com/go-ego/gse/types"
)

var (
	// ToLower set alpha to lowercase
	ToLower = true
)

const (
	zhS1 = "dict/zh/s_1.txt"
	zhT1 = "dict/zh/t_1.txt"
)

// Init initializes the segmenter config
func (seg *Segmenter) Init() { _ = "STUB: not implemented"; return }

// init the model of hmm cut

// Dictionary returns the dictionary used by the tokenizer
func (seg *Segmenter) Dictionary() *Dictionary {
	_ = "STUB: not implemented"

	// ToToken make the text, freq and pos to token structure
	return nil
}

func (seg *Segmenter) ToToken(text string, freq float64, pos ...string) Token {
	_ = "STUB: not implemented"
	return *new(Token)
}

// AddToken add a new text to the token
func (seg *Segmenter) AddToken(text string, freq float64, pos ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddTokenForce add new text to token and force
// time-consuming
func (seg *Segmenter) AddTokenForce(text string, freq float64, pos ...string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ReAddToken remove and add token again
func (seg *Segmenter) ReAddToken(text string, freq float64, pos ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveToken remove token in dictionary
func (seg *Segmenter) RemoveToken(text string) error { _ = "STUB: not implemented"; return nil }

// Empty empty the seg dictionary
func (seg *Segmenter) Empty() error { _ = "STUB: not implemented"; return nil }

// LoadDictMap load dictionary from []map[string]string
func (seg *Segmenter) LoadDictMap(dict []map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse the word frequency

// LoadDict load the dictionary from the file
//
// The format of the dictionary is (one for each participle):
//
//	participle text, frequency, part of speech
//
// # And you can option the dictionary separator by seg.DictSep = ","
//
// Can load multiple dictionary files, the file name separated by "," or ", "
// the front of the dictionary preferentially load the participle,
//
//	such as: "user_dictionary.txt,common_dictionary.txt"
//
// When a participle appears both in the user dictionary and
// in the `common dictionary`, the `user dictionary` is given priority.
func (seg *Segmenter) LoadDict(files ...string) error { _ = "STUB: not implemented"; return nil }

// seg.Load = true

// load     bool

// return errors.New("Dict files is nil.")

// load = true
// files = dictFiles

// files = []string{dictPath}

// if files[0] != "" && files[0] != "en" && !load {
// 	for _, file := range strings.Split(files[0], ",") {
// 		// for _, file := range files {
// 		err := seg.Read(file)
// 		if err != nil {
// 			return err
// 		}
// 	}
// }

// LoadTFIDFDict load tfidf dict for cal tfidf & bm25
func (seg *Segmenter) LoadTFIDFDict(files []*types.LoadDictFile) error {
	_ = "STUB: not implemented"
	return nil
}

// seg.Load = true

// return errors.New("Dict files is nil.")

// GetCurrentFilePath get the current file path
func (seg *Segmenter) GetCurrentFilePath() string { _ = "STUB: not implemented"; return "" }

// GetIdfPath get the idf path
func (seg *Segmenter) GetIdfPath(files ...string) []string { _ = "STUB: not implemented"; return nil }

// LoadCorpusAverLen load the average length of corpus
func (seg *Segmenter) LoadCorpusAverLen(files ...string) (corpusTotal float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetCorpusPath get the corpus path
func (seg *Segmenter) GetCorpusPath(files ...string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (seg *Segmenter) ReadCorpus(file string) (corpusAverLen float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// new the Scanner to read file content

// read file content by line

// GetTfIdfPath get the tfidf path
func (seg *Segmenter) GetTfIdfPath(files ...string) []string { _ = "STUB: not implemented"; return nil }

// Read read the dict file
func (seg *Segmenter) Read(file string) error { _ = "STUB: not implemented"; return nil }

// ReadTFIDF read the dict file
func (seg *Segmenter) ReadTFIDF(file string) error { _ = "STUB: not implemented"; return nil }

// Size frequency is calculated based on the size of the text
func (seg *Segmenter) Size(size int, text, freqText string) (freq float64) {
	_ = "STUB: not implemented"

	// End of file or error line
	// continue
	return 0
}

// invalid row line

// Analyze the word frequency

// continue

// Filter out the words that are too infrequent

// Filter words with a length of 1 to reduce the word frequency

// ReadN read the tokens by '\n'
func (seg *Segmenter) ReadN(reader *bufio.Reader) (size int, text, freqText, pos string, fsErr error) {
	_ = "STUB: not implemented"
	return 0, "", "", "", nil
}

// ReadNTFIDF read the tokens with tfidf by '\n'
func (seg *Segmenter) ReadNTFIDF(reader *bufio.Reader) (size int, text, freqText, idfText string, fsErr error) {
	_ = "STUB: not implemented"
	return 0, "", "", "", nil
}

// Reader load dictionary from io.Reader
func (seg *Segmenter) Reader(reader *bufio.Reader, files ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Read the word segmentation line by line

// End of file

// No part of speech, marked as an empty string

// Add participle tokens to the dictionary

// ReaderTFIDF load tfidf dictionary from io.Reader
func (seg *Segmenter) ReaderTFIDF(reader *bufio.Reader, files ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Read the word segmentation line by line

// End of file

// Add participle tokens to the dictionary

// DictPaths get the dict's paths
func DictPaths(dictDir, filePath string) (files []string) { _ = "STUB: not implemented"; return nil }

// if str[i] == "ti" {
// }

// }

// IsJp is Japan char return true
func IsJp(segText string) bool { _ = "STUB: not implemented"; return false }

// CalcToken calc the segmenter token
func (seg *Segmenter) CalcToken() {
	_ = "STUB: not implemented"
	// Calculate the path value of each word segment.
	// For the meaning of the path value, see the notes of the Token structure
	return
}

// Each word segmentation is carefully divided for search engine mode,
// For the usage of this mode, see the comments of the Token structure.

// Calculate the number of sub-segments that need to be added

// add sub-segments subparticiple
