// Package src provides utility functions for demonstration.
//
// Author: Sefat Anam (sefatanam@gmail.com)
// Created: 2025-10-10 00:26:34 +0600 (+06)
package src

func BulkProcess(r int, c chan int) {
	sum := 0
	for i := range r {
		c <- i
		sum += i
	}
	c <- sum
	close(c)
}


