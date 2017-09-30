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

// Pixtrix is the type fore representing Pixel Matrix
type Pixtrix struct {
	Rows   int      `json:"rows"`
	Bitmap []uint32 `json:"bitmap"`
}

// NewMatrix creates a new pixel matrix with the givent size. The rows
// is fixed and the columns can be expanded if needed.
func NewMatrix(rows, columns int) *Pixtrix {
	m := new(Pixtrix)
	m.Rows = rows
	m.Bitmap = make([]uint32, rows*columns)
	return m
}

// Columns returns the current number of columns of the pixel matrix.
func (m *Pixtrix) Columns() int {
	if len(m.Bitmap)%m.Rows != 0 {
		panic("invalid matrix length")
	}
	return len(m.Bitmap) / m.Rows
}

// CheckAndResize checks if the coordinate (x,y) is valid in the Matrix. Extend the
// Matrix if there is not enough space for x.
func (m *Pixtrix) CheckAndResize(x, y int) {
	if y < 0 || y >= m.Rows {
		panic("y out of bound")
	}
	if x < 0 {
		panic("x out of bound")
	}
	if x >= m.Columns() { // resize
		requiredSize := (x + 1) * m.Rows
		if cap(m.Bitmap) < requiredSize {
			t := make([]uint32, requiredSize)
			copy(t, m.Bitmap)
			m.Bitmap = t
		} else {
			m.Bitmap = m.Bitmap[:requiredSize]
		}
	}
}

// SetPixel paints the pixel at column "x" and row "y" with the given color.
func (m *Pixtrix) SetPixel(x, y int, color uint32) {
	m.CheckAndResize(x, y)
	m.Bitmap[x*m.Rows+y] = color
}

// SetPixelAlpha paints the pixel at column "x" and row "y" with the given color and
// the given alpha channel between 0 (transparent) and 255 (opaque).
func (m *Pixtrix) SetPixelAlpha(x, y int, color uint32, alpha int) {
	m.CheckAndResize(x, y)
	r0, g0, b0 := DeRGB(m.Bitmap[x*m.Rows+y])
	r1, g1, b1 := DeRGB(color)
	r := (r0*alpha + r1*(255-alpha)) / 255
	g := (g0*alpha + g1*(255-alpha)) / 255
	b := (b0*alpha + b1*(255-alpha)) / 255
	m.Bitmap[x*m.Rows+y] = RGB(r, g, b)
}

// GetPixel returns the color ot the pixel at column "x" and row "y".
func (m *Pixtrix) GetPixel(x, y int) uint32 {
	if y < 0 || y >= m.Rows {
		panic("y out of bound")
	}
	if x < 0 || x >= m.Columns() {
		panic("x out of bound")
	}
	return m.Bitmap[x*m.Rows+y]
}

// Slice cuts a pixel matrix at 2 vertical locations (low and high) and returns
// the resulting pixel materix.
func (m *Pixtrix) Slice(low, high int) *Pixtrix {
	res := new(Pixtrix)
	res.Bitmap = m.Bitmap[low*m.Rows : high*m.Rows]
	res.Rows = m.Rows
	return res
}

// Append append other pixel matrices to the current one.
func (m *Pixtrix) Append(x ...*Pixtrix) {
	for _, i := range x {
		if m.Rows != i.Rows {
			panic("Error Matrix Operation")
		}
		m.Bitmap = append(m.Bitmap, i.Bitmap...)
	}
}

// Concat is similar to Append, but it does not change the current
// pixel matrix. Instead it returns a new one.
func Concat(a *Pixtrix, b ...*Pixtrix) *Pixtrix {
	res := new(Pixtrix)
	res.Rows = a.Rows
	res.Bitmap = append(res.Bitmap, a.Bitmap...)
	res.Append(b...)
	return res
}

// InterleavedStripes convert the bitmap according to the constructions of most
// LED matrix devices, where the even and the odd columns have not the same order
// of pixels. In other terms, The pixels are "zigzagging" on every column.
// The result of the method is 2 array. The first array (even) is an
// array of uint32 that diplays correctly at even locations (0, 2, 4, ...). The
// second arrays displays correctly at odd locations (1, 3, 5, ...).
func (m *Pixtrix) InterleavedStripes() (even []uint32, odd []uint32) {
	even = make([]uint32, m.Rows*m.Columns())
	odd = make([]uint32, m.Rows*m.Columns())
	for x := 0; x < m.Columns(); x++ {
		for y := 0; y < m.Rows; y++ {
			if x%2 == 0 { // Even Column
				even[x*m.Rows+y] = m.Bitmap[x*m.Rows+y]
				odd[x*m.Rows+y] = m.Bitmap[x*m.Rows+m.Rows-1-y]
			} else { // Odd Column
				even[x*m.Rows+y] = m.Bitmap[x*m.Rows+m.Rows-1-y]
				odd[x*m.Rows+y] = m.Bitmap[x*m.Rows+y]
			}
		}
	}
	return
}

// StripeToBytes extracts the RBG components of the stripe and returns a corresponding
// array of bytes.
func StripeToBytes(stripe []uint32) []byte {
	res := make([]byte, len(stripe)*3)
	for i, v := range stripe {
		res[i*3] = byte(v >> 16 & 0xff)
		res[i*3+1] = byte(v >> 8 & 0xff)
		res[i*3+2] = byte(v & 0xff)
	}
	return res
}
