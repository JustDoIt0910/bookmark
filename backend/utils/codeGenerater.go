package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func NewCode() string {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}
