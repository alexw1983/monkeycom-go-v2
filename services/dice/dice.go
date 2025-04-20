package dice

import (
	"log"
	"math/rand/v2"
)

func RollDX(x int) int {
	log.Printf("RollD%d", x)

	var val = rand.IntN(x-1) + 1

	log.Printf("RollD%d = %d", x, val)

	return val
}
