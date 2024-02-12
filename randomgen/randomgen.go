package randomgen

import (
	"math"
	"math/rand"
	"time"
)

func RandomInt(min int, max int) int {
	//Generate a random integer with the range given
	//Accomodate negative values as well

	randBase := rand.New(rand.NewSource(time.Now().UnixNano()))

	var modifiedMin int = 0
	var modifiedMax int = 0

	//Adjust max for negative values if needed
	//--------------------------------------------
	if min < 0 {
		modifiedMax = max - min
	} else {
		modifiedMax = max
	}

	//Adjust min for negative values if needed
	//--------------------------------------------
	if max < 0 {
		modifiedMin = max * -1
	} else {
		modifiedMin = min
	}

	//Generate random number
	var randNum int = 0

	if min < 0 {
		randNum = (randBase.Intn(modifiedMax - (modifiedMin + 1) + modifiedMin)) + min
	} else {
		randNum = randBase.Intn(max - (min + 1) + min)
	}

	return randNum

}

func RandomIFloat(min float64, max float64, scale int) float64 {
	//Generate a random float value with the range and scale given
	//Accomodate negative values as well

	randBase := rand.New(rand.NewSource(time.Now().UnixNano()))

	//Adjust inputs to appropriate scale
	//fmt.Println(min * math.Pow10(scale))

	var minInt int = int(min * math.Pow10(scale))
	var maxInt int = int(max * math.Pow10(scale))

	var modifiedMin int = 0
	var modifiedMax int = 0

	//Adjust max for negative values if needed
	//--------------------------------------------
	if minInt < 0 {
		modifiedMax = maxInt - minInt
	} else {
		modifiedMax = maxInt
	}

	//Adjust min for negative values if needed
	//--------------------------------------------
	if maxInt < 0 {
		modifiedMin = maxInt * -1
	} else {
		modifiedMin = minInt
	}

	//Generate random number
	var randNum int = 0
	var randBoundary int = 0

	//fmt.Println(min, max, scale, minInt, maxInt, modifiedMin, modifiedMax)

	if min < 0 {
		randBoundary = modifiedMax - (modifiedMin + 1)
		//fmt.Println(randBoundary)
		randNum = (randBase.Intn(randBoundary) + modifiedMin) + minInt
	} else {
		randBoundary = modifiedMax - (modifiedMin + 1)
		//fmt.Println(randBoundary)
		randNum = randBase.Intn(randBoundary) + modifiedMin
	}

	return float64(randNum) / math.Pow10(scale)

}

func RandomBool() int {
	//Generate a random 1 or 0 (binary boolean)
	randBase := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randBase.Intn(2)
}
