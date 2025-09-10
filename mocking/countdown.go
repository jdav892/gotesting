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

type SpyCountdownOperations struct {
	Calls []string
}

//add new line character to go string for correct terminal output with this function
const finalWord = "Go!"
const countDownStart = 3
const write = "write"
const sleep = "sleep"

func Countdown(out io.Writer, sleeper Sleeper){
	//using for loop to decrement and recreate a countdown
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}


func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

