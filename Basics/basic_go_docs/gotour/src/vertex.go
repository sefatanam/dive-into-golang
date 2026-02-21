// Package src provides utility functions for demonstration.
//
// Author: Sefat Anam (sefatanam@gmail.com)
// Created: 2025-10-09 22:31:56 +0600 (+06)
package src

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// func (r *Rectangle) Scale(f float64) {
// 	r.Height = r.Height * f
// 	r.Width = r.Width * f
// }

func Scale(r Rectangle,f float64) {
	r.Height = r.Height * f
	r.Width = r.Width * f
}

