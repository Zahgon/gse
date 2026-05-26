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

type probState struct {
	prob  float64
	state uint16
}

func (ps probState) String() string { _ = "STUB: not implemented"; return "" }

type probStates []probState

func (pss probStates) Len() int { _ = "STUB: not implemented"; return 0 }

func (pss probStates) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pss probStates) Swap(i, j int) { _ = "STUB: not implemented"; return }

func viterbi(obs []rune) []tag { _ = "STUB: not implemented"; return nil }

// default is all_states

func probs(obs []rune, vtb []map[uint16]float64, memPath []map[uint16]uint16,
	obsLen int) ([]map[uint16]uint16, []map[uint16]float64) {
	_ = "STUB: not implemented"
	return nil, nil
}
