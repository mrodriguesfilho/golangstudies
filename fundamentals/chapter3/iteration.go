package iteration

func Repeat(character string, amountOfTimes int) string {

	var repeated string

	for i := 0; i < amountOfTimes; i++ {
		repeated += character
	}

	return repeated
}
