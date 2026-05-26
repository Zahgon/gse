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
	"github.com/go-ego/gse"
)

// Dict type a dictionary for the word segmentation.
type Dict struct {
	total, logTotal float64
	Seg             gse.Segmenter
}

// func New(files ...string) gse.Segmenter {
// 	return gse.New(files...)
// }

// AddToken add new text to token
func (d *Dict) AddToken(text string, freq float64, pos ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveToken remove the token in the dict
func (d *Dict) RemoveToken(text string) error { _ = "STUB: not implemented"; return nil }

func (d *Dict) updateLogTotal() { _ = "STUB: not implemented"; return }

// Freq find the word return the word's freq, pos and existence
func (d *Dict) Freq(key string) (float64, string, bool) {
	_ = "STUB: not implemented"
	return 0,

		// Pos find the key return the POS and existence
		"", false
}

func (d *Dict) Pos(key string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (d *Dict) loadDict(files ...string) error { _ = "STUB: not implemented"; return nil }
