// Copyright 2017 Jacques Supcik, Blue Masters
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pixtrix

import (
	"testing"
)

func TestBasic(t *testing.T) {
	m := NewMatrix(8, 0)
	m.SetPixel(4, 4, White)
	if m.Columns() != 5 {
		t.Errorf("Columns shoud now be 5 instead of %d", m.Rows)
	}
	m.SetPixel(2, 2, Red)
	if m.Columns() != 5 {
		t.Errorf("Columns shoud still be 5 instead of %d", m.Rows)
	}

	if m.GetPixel(2, 2) != Red {
		t.Error("Pixel shoud be red")
	}
	if m.GetPixel(3, 3) != Black {
		t.Error("Pixel shoud be black")
	}
	if m.GetPixel(4, 4) != White {
		t.Error("Pixel shoud be white")
	}
}

func TestSlice(t *testing.T) {
	m := NewMatrix(8, 10)
	m.SetPixel(4, 4, White)
	m.SetPixel(5, 4, Red)
	m.SetPixel(6, 4, Red)
	m.SetPixel(7, 4, White)

	n := m.Slice(4, 8)
	if n.Rows != m.Rows {
		t.Error("Number of rows is not the same")
	}
	if n.Columns() != 4 {
		t.Errorf("Number of columns is %d instead of 4", n.Columns())
	}

	if n.GetPixel(0, 4) != White {
		t.Error("Pixel shoud be white")
	}
	if n.GetPixel(1, 4) != Red {
		t.Error("Pixel shoud be red")
	}
	if n.GetPixel(2, 4) != Red {
		t.Error("Pixel shoud be red")
	}
	if n.GetPixel(3, 4) != White {
		t.Error("Pixel shoud be white")
	}

	if n.GetPixel(0, 0) != Black {
		t.Error("Pixel shoud be black")
	}
	if n.GetPixel(3, 5) != Black {
		t.Error("Pixel shoud be black")
	}

}
