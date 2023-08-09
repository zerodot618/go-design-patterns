package factory

import (
	"fmt"
	"testing"
)

func TestFactorty(t *testing.T) {
	ak47, _ := getGun("ak47")
	maverick, _ := getGun("maverick")
	printlnDetails(ak47)
	printlnDetails(maverick)
}

func printlnDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
