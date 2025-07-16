package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3


func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}

type Sleeper interface{
	Sleep()
}

// Mocked sleeper
type SpySleeper struct{
	Calls int
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

// Actuall sleeper
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep(){
	time.Sleep(1 *time.Second)
}


// Another dependency
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"



type ConfigurableSleeper struct{
	duration time.Duration
	sleep func(time.Duration)
}


type SpyTime struct{
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration){
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep(){
	c.sleep(c.duration)
}
