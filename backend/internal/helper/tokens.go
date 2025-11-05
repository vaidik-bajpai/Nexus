package helper

import (
	"math/rand"
	"time"
)

func CreateRandomToken() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(1000000)
}
