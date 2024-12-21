package iteration

func Repeat(char string, times int) string {
	var output string
	for range times {
		output += char
	}
	return output
}
