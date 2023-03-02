package iteration

func Repeat(character string, repeatTimes int) string {
	var repeated string
	for i := 0; i < repeatTimes; i++ {
		repeated += character
	}
	return repeated
}
