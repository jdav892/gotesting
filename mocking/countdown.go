package main

import(
	"fmt"
	"io"
	"time"
	"os"
)

const finalWord = "Go!"
const countDown = 3

func Countdown(out io.Writer){
	//using for loop to decrement and recreate a countdown
	for i:= countDown; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}


func main() {
	Countdown(os.Stdout)
}

