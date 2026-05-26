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

package hmm

const minFloat = -3.14e100

var (
	prevStatus = make(map[byte][]byte)
	probStart  = make(map[byte]float64)
)

func init() {
	prevStatus['B'] = []byte{'E', 'S'}
	prevStatus['M'] = []byte{'M', 'B'}
	prevStatus['S'] = []byte{'S', 'E'}
	prevStatus['E'] = []byte{'B', 'M'}

	probStart['B'] = -0.26268660809250016
	probStart['E'] = -3.14e+100
	probStart['M'] = -3.14e+100
	probStart['S'] = -1.4652633398537678
}

type probState struct {
	prob  float64
	state byte
}

func (p probState) String() string { _ = "STUB: not implemented"; return "" }

type probStates []*probState

func (ps probStates) Len() int { _ = "STUB: not implemented"; return 0 }

func (ps probStates) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (ps probStates) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Viterbi write viterbi algorithm by go
func Viterbi(obs []rune, states []byte) (float64, []byte) { _ = "STUB: not implemented"; return 0, nil }

// path, newPath := paths(obs, states, vtb, path, t)

func probs(obs []rune, vtb []map[byte]float64, y byte, t int) (ps0 probStates) {
	_ = "STUB: not implemented"
	return *new(probStates)
}
