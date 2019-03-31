package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	randString := "AUOPNHYJMLQBEVCRXZMK"

	rand.Seed(time.Now().Unix())
	newInt := rand.Intn(1000)

	for i := 0; i < 10; i++ {
		thisMoment := time.Now().UnixNano()
		fmt.Println(thisMoment)
	}

}
