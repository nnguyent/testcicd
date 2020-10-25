package SilentiumIO

import (
	"fmt"
)

// SpecNum -
type SpecNum int

const (
	// NUM4 -
	NUM4 SpecNum = iota
	// NUM7 -
	NUM7
	// NUM28 -
	NUM28
)

// Value -
func (sn SpecNum) Value() int {
	val := 0
	switch sn {
	case NUM4:
		val = 4
	case NUM7:
		val = 7
	case NUM28:
		val = 28
	}

	return val
}

// String -
func (sn SpecNum) String() string {
	return [...]string{
		"Silentium",
		"IO",
		"SilentiumIO"}[sn]
}

// PrintNumber -
func PrintNumber(n int) {
	for i := 1; i <= n; i++ {
		if i%NUM28.Value() == 0 {
			fmt.Println(NUM28.String())
		} else if i%NUM4.Value() == 0 {
			fmt.Println(NUM4.String())
		} else if i%NUM7.Value() == 0 {
			fmt.Println(NUM7.String())
		} else {
			fmt.Println(i)
		}
	}
}
