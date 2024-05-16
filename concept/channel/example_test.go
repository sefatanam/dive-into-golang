package channel

import "testing"

func TestExample(t *testing.T) {
	numbers := NewNumbers()
	numbers.Start()

	for i := 1; i <= 10; i++ {
		/* go func(i int) {
			numbers.numsChan <- i
		}(i) */
		numbers.NumsChan <- i
	}
}
