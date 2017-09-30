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
	"github.com/telecom-tower/font"
)

// Writer is the type for writing texts and bitmaps to pixel matrix.
type Writer struct {
	matrix *Pixtrix
	Pos    int // Cursor position
}

// NewWriter creates a new writer on the given pixel matrix.
func NewWriter(m *Pixtrix) *Writer {
	w := new(Writer)
	w.matrix = m
	w.Pos = 0
	return w
}

// WriteTextAlpha writes the given text with the given font, color and alpha channel (transparency)
func (w *Writer) WriteTextAlpha(text string, fnt font.Font, color uint32, alpha int) {
	for _, c := range font.ExpandAlias(text) {
		for _, i := range fnt.Bitmap[c] {
			for k := 0; k < fnt.Height; k++ {
				if i&(1<<uint(k)) != 0 {
					w.matrix.SetPixelAlpha(w.Pos, k, color, alpha)
				}
			}
			w.Pos++
		}
	}
}

// WriteText writes the given text with the given font, color and backfround color
func (w *Writer) WriteText(text string, fnt font.Font, color, bgColor uint32) {
	for _, c := range font.ExpandAlias(text) {
		for _, i := range fnt.Bitmap[c] {
			for k := 0; k < fnt.Height; k++ {
				if i&(1<<uint(k)) != 0 {
					w.matrix.SetPixel(w.Pos, k, color)
				} else {
					w.matrix.SetPixel(w.Pos, k, bgColor)
				}
			}
			w.Pos++
		}
	}
}

// WriteBitmap writes a bitmap into the pixel matrix through the writer. This is usually
// used to insert PNG images into the text (e.g icon, smiley, ...)
func (w *Writer) WriteBitmap(bitmap [][]uint32) {
	width := 0
	for y := 0; y < len(bitmap); y++ {
		if len(bitmap[y]) > width {
			width = len(bitmap[y])
		}
		for x := 0; x < len(bitmap[y]); x++ {
			w.matrix.SetPixel(w.Pos+x, y, bitmap[y][x])
		}
	}
	w.Pos += width
}

// Spacer insert a space with the given width and color
func (w *Writer) Spacer(width int, color uint32) {
	for y := 0; y < w.matrix.Rows; y++ {
		for x := 0; x < width; x++ {
			w.matrix.SetPixel(w.Pos+x, y, color)
		}
	}
	w.Pos += width
}
