package detail

import (
	"fmt"
	"math/rand"
)

func Showinfo() {
	fmt.Println("dsadsafdsf")
}
func GenRound() string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randoms := make([]rune, 5)
	for i := 0; i < 5; i++ {
		randoms[i] = letters[rand.Intn(len(letters))]
	}
	return string(randoms)
}
