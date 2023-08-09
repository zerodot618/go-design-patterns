package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton2(t *testing.T) {
	for i := 0; i < 100; i++ {
		go GetInstance2()
	}
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
