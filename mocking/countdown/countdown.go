package countdown

import (
	"fmt"
	"io"
	"strconv"
)

func Countdown(writer io.Writer, count int, final string) {
	for i := count; i > 0; i-- {
		fmt.Fprint(writer, strconv.Itoa(i))
	}
	fmt.Fprint(writer, final)
}
