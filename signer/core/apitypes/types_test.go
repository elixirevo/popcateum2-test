// Copyright 2023 The go-popcateum Authors
// This file is part of the go-popcateum library.
//
// The go-popcateum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-popcateum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-popcateum library. If not, see <http://www.gnu.org/licenses/>.

package apitypes

import "testing"

func TestIsPrimitive(t *testing.T) {
	// Expected positives
	for i, tc := range []string{
		"int24", "int24[]", "uint88", "uint88[]", "uint", "uint[]", "int256", "int256[]",
		"uint96", "uint96[]", "int96", "int96[]", "bytes17[]", "bytes17",
	} {
		if !isPrimitiveTypeValid(tc) {
			t.Errorf("test %d: expected '%v' to be a valid primitive", i, tc)
		}
	}
	// Expected negatives
	for i, tc := range []string{
		"int257", "int257[]", "uint88 ", "uint88 []", "uint257", "uint-1[]",
		"uint0", "uint0[]", "int95", "int95[]", "uint1", "uint1[]", "bytes33[]", "bytess",
	} {
		if isPrimitiveTypeValid(tc) {
			t.Errorf("test %d: expected '%v' to not be a valid primitive", i, tc)
		}
	}
}
