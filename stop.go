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

// StopWordMap the default stop words.
var StopWordMap = map[string]bool{
	" ": true,
}

// LoadStopArr load stop word by []string
func (seg *Segmenter) LoadStopArr(dict []string) { _ = "STUB: not implemented"; return }

// LoadStop load stop word files add token to map
func (seg *Segmenter) LoadStop(files ...string) error { _ = "STUB: not implemented"; return nil }

// AddStop add a token to the StopWord dictionary.
func (seg *Segmenter) AddStop(text string) { _ = "STUB: not implemented"; return }

// AddStopArr add array stop token to stop dictionaries
func (seg *Segmenter) AddStopArr(text ...string) { _ = "STUB: not implemented"; return }

// RemoveStop remove a token from the StopWord dictionary.
func (seg *Segmenter) RemoveStop(text string) { _ = "STUB: not implemented"; return }

// EmptyStop empty the stop dictionary
func (seg *Segmenter) EmptyStop() error { _ = "STUB: not implemented"; return nil }

// IsStop check the word is a stop word.
func (seg *Segmenter) IsStop(s string) bool { _ = "STUB: not implemented"; return false }

// return StopWordMap[s]
