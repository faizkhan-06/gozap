package utils

import (
	"math/big"
	"math/rand"
	"time"

	"github.com/mattheath/base62"
)

func GenerateShortId() string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	num := r.Uint64() % 56800235584
	return base62.EncodeBigInt(big.NewInt(int64(num)))
}