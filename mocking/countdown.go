package main

import(
	"fmt"
	"io"
	"time"
	"os"
)

type DefaultSleeper struct {

}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

//add new line character to go string for correct terminal output with this function
const finalWord = "Go!"
const countDownStart = 3

func Countdown(out io.Writer, sleeper Sleeper){
	//using for loop to decrement and recreate a countdown
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}


func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

