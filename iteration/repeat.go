package iteration

// Repeat takes a string & an integer and return the string repeated by the value of the integer
func Repeat(character string, repeat int) string {
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += character
	}
	return repeated
}
