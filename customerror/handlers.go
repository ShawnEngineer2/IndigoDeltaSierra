package customerror

func CheckAndPanic(e error) {
	if e != nil {
		panic(e)
	}
}
