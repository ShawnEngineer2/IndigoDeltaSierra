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

func CreateDistributionMatrix() []int {
	//Returns a matrix of numbers that can be used as part of a sampling distribution
	distro := make([]int, 139)

	distro[0] = 0
	distro[1] = 0
	distro[2] = 0
	distro[3] = 0
	distro[4] = 0
	distro[5] = 0
	distro[6] = 1
	distro[7] = 0
	distro[8] = 0
	distro[9] = 0
	distro[10] = 0
	distro[11] = 0
	distro[12] = 0
	distro[13] = 2
	distro[14] = 2
	distro[15] = 2
	distro[16] = 0
	distro[17] = 0
	distro[18] = 0
	distro[19] = 0
	distro[20] = 0
	distro[21] = 0
	distro[22] = 3
	distro[23] = 3
	distro[24] = 3
	distro[25] = 3
	distro[26] = 0
	distro[27] = 0
	distro[28] = 0
	distro[29] = 0
	distro[30] = 0
	distro[31] = 0
	distro[32] = 2
	distro[33] = 2
	distro[34] = 2
	distro[35] = 0
	distro[36] = 0
	distro[37] = 0
	distro[38] = 0
	distro[39] = 0
	distro[40] = 0
	distro[41] = 1
	distro[42] = 0
	distro[43] = 0
	distro[44] = 0
	distro[45] = 0
	distro[46] = 0
	distro[47] = 0
	distro[48] = 1
	distro[49] = 0
	distro[50] = 0
	distro[51] = 0
	distro[52] = 0
	distro[53] = 0
	distro[54] = 0
	distro[55] = 2
	distro[56] = 2
	distro[57] = 2
	distro[58] = 0
	distro[59] = 0
	distro[60] = 0
	distro[61] = 0
	distro[62] = 0
	distro[63] = 0
	distro[64] = 3
	distro[65] = 3
	distro[66] = 3
	distro[67] = 3
	distro[68] = 0
	distro[69] = 0
	distro[70] = 0
	distro[71] = 0
	distro[72] = 0
	distro[73] = 0
	distro[74] = 2
	distro[75] = 2
	distro[76] = 2
	distro[77] = 0
	distro[78] = 0
	distro[79] = 0
	distro[80] = 0
	distro[81] = 0
	distro[82] = 0
	distro[83] = 1
	distro[84] = 0
	distro[85] = 0
	distro[86] = 0
	distro[87] = 0
	distro[88] = 0
	distro[89] = 0
	distro[90] = 3
	distro[91] = 0
	distro[92] = 3
	distro[93] = 0
	distro[94] = 2
	distro[95] = 0
	distro[96] = 1
	distro[97] = 1
	distro[98] = 0
	distro[99] = 0
	distro[100] = 0
	distro[101] = 0
	distro[102] = 0
	distro[103] = 0
	distro[104] = 2
	distro[105] = 2
	distro[106] = 2
	distro[107] = 0
	distro[108] = 0
	distro[109] = 0
	distro[110] = 0
	distro[111] = 0
	distro[112] = 0
	distro[113] = 3
	distro[114] = 3
	distro[115] = 3
	distro[116] = 3
	distro[117] = 0
	distro[118] = 0
	distro[119] = 0
	distro[120] = 0
	distro[121] = 0
	distro[122] = 0
	distro[123] = 2
	distro[124] = 2
	distro[125] = 2
	distro[126] = 0
	distro[127] = 0
	distro[128] = 0
	distro[129] = 0
	distro[130] = 0
	distro[131] = 0
	distro[132] = 1
	distro[133] = 0
	distro[134] = 0
	distro[135] = 0
	distro[136] = 0
	distro[137] = 0
	distro[138] = 0

	return distro

}
