package cache

import (
	"fmt"
	"testing"
)

func TestCalcLenest_cache(t *testing.T) {
	v := CalcLen(1)
	fmt.Println(v)
}
