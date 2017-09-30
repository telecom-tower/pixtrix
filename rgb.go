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

// RGB creates a color based of the red, green and blue components.
func RGB(r, g, b int) uint32 {
	if r < 0 || r > 255 {
		panic("Red component must be between 0 and 255")
	}
	if g < 0 || g > 255 {
		panic("Green component must be between 0 and 255")
	}
	if b < 0 || b > 255 {
		panic("Blue component must be between 0 and 255")
	}
	return uint32(r)<<16 + uint32(g)<<8 + uint32(b)
}

// DeRGB decodes a color and returns the red, green and blue components.
func DeRGB(color uint32) (r, g, b int) {
	b = int(color & 0xff)
	g = int((color >> 8) & 0xff)
	r = int((color >> 16) & 0xff)
	return
}

// Red is the red color (#FF0000)
var Red = RGB(255, 0, 0)

// Green is the green color (#00FF00)
var Green = RGB(0, 255, 0)

// Blue is the blue color (#0000FF)
var Blue = RGB(0, 0, 255)

// White is the white color (#FFFFFF)
var White = RGB(255, 255, 255)

// Black is no color (#000000)
var Black = RGB(0, 0, 0)
