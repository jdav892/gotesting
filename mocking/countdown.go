package main

import(
	"fmt"
	"io"
	"time"
	"os"
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

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

type SpyTime struct {
	durationSlept time.Duration
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

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}


func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

