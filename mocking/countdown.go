package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 3 to out.
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	fmt.Fprintf(out, finalWord)

}

func main() {
	Countdown(os.Stdout)
}
