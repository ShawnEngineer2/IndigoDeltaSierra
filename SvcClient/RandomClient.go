package SvcClient

import (
	"fmt"
)

func GetRandomNumbers() {

	svcUrl := "https://www.random.org/integers/?num=5&min=1&max=6&col=1&base=10&format=plain&rnd=new"

	responseBytes := Get(svcUrl, "", nil)
	responseString := string(responseBytes)

	fmt.Println(responseString)

}
