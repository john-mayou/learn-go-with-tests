package countdown

import (
	"fmt"
	"io"
	"strconv"
)

type Sleeper interface {
	Sleep()
}

func Countdown(writer io.Writer, sleeper Sleeper, count int, final string) {
	for i := count; i > 0; i-- {
		fmt.Fprintln(writer, strconv.Itoa(i))
		sleeper.Sleep()
	}
	fmt.Fprint(writer, final)
}
