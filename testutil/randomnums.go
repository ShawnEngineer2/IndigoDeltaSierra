package testutil

import (
	"fmt"
	"indigodeltasierra/randomgen"
	"math/rand"
	"time"
)

func TestRandomNumbers() {

	var min int = 0
	var negmin int = -230
	var max int = 200
	var scale int = 5
	var randNum int = 1
	var randNumFloat float64 = 0.00

	randBase := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("Just a random number")
	randNum = randBase.Int()
	fmt.Println(randNum)

	fmt.Println("Constrained Random number (0 to 65)")
	randNum = randBase.Intn(max + 35)
	fmt.Println(randNum)

	fmt.Println("Constrained Random number range (10 to 30)")
	randNum = randBase.Intn(max - (min + 1) + min)
	fmt.Println(randNum)

	fmt.Println("Scaled Number to represent Float operations")
	randNum = randBase.Intn((max * scale) - ((min * scale) + 1) + (min * scale))
	randNumFloat = float64(randNum) / float64(scale)
	fmt.Println(randNumFloat)

	fmt.Println("Constrained Random negative number range (10 to 30)")

	if negmin < 0 {
		min = max
		max = max - negmin
	}

	randNum = (randBase.Intn(max - (min + 1) + min)) + negmin
	fmt.Println(randNum)

	fmt.Println("Constrained Random boolean")
	randNum = randBase.Intn(2)
	fmt.Println(randNum)

}

func TestRandomNumbers02() {

	var randNumInt int = 0
	var randNumFloat float64 = 0.00

	fmt.Println("Test range of positive integers (5 - 12)")
	randNumInt = randomgen.RandomInt(5, 12)
	fmt.Println(randNumInt)

	fmt.Println("Test range of positive float (.0032 - 100.00)")
	randNumFloat = randomgen.RandomIFloat(.0032, 100.00, 5)
	fmt.Println(randNumFloat)

	fmt.Println("Test Binary Boolean")
	randNumInt = randomgen.RandomBool()
	fmt.Println(randNumInt)

	fmt.Println("Test range of mixed integers (-5 - 7)")
	randNumInt = randomgen.RandomInt(-5, 7)
	fmt.Println(randNumInt)

	fmt.Println("Test range of mixed floats (-11.82738 - 4.55000)")
	randNumFloat = randomgen.RandomIFloat(-11.82738, 4.55000, 5)
	fmt.Println(randNumFloat)

	fmt.Println("Test range of negative integers (-22 - -11)")
	randNumInt = randomgen.RandomInt(-22, -11)
	fmt.Println(randNumInt)

	fmt.Println("Test range of negative floats (-11.82738 - -4.55000)")
	randNumFloat = randomgen.RandomIFloat(-11.82738, -4.55000, 5)
	fmt.Println(randNumFloat)

}
