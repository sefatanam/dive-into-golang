// Package src 
//
// Author: Sefat Anam (sefatanam@gmail.com)
// Created: 2025-10-09 22:21:14 +0600 (+06)
//
package src

type MyFloat float64

func (mf MyFloat) MakeUint() uint {
	return uint(mf)
}
