package iteration

func Repeat(char string) string {
	var output string
	for range 5 {
		output += char
	}
	return output
}
