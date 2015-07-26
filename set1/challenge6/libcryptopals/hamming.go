package libcryptopals

import (
	"errors"
	"fmt"
)

func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New(fmt.Sprintf("Input strings must have equal length, not ", len(a), len(b)))
	}
	var distance int = 0
	for index, val := range a {
		// XOR == light up all the bit differences
		xord := val ^ b[index]
		// I suspect there's some bitwise magic you could use to do this, but I'm not feeling super ambitious
		for i := uint(0); i < 8; i++ {
			if (1<<i)&xord != 0 {
				distance++
			}
		}
	}
	return distance, nil
}
