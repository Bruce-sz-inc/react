package timer //import "myitcv.io/react/examples/timer"

import (
	"fmt"
	"time"

	r "myitcv.io/react"
)

//go:generate reactGen

// TimerDef is the definition of the Timer component
type TimerDef struct {
	r.ComponentDef
}

// TimerState is the state type for the Timer component
type TimerState struct {
	ticker *time.Ticker

	secondsElapsed float32
}

// Timer creates instances of the Timer component
func Timer() *TimerDef {
	res := new(TimerDef)
	r.BlessElement(res, nil)
	return res
}

// ComponentWillMount is a React lifecycle method for the Timer component
func (t *TimerDef) ComponentWillMount() {
	tick := time.NewTicker(time.Second * 1)

	s := t.State()
	s.ticker = tick
	t.SetState(s)

	go func() {
		for {
			select {
			case <-tick.C:
				c := t.State()
				c.secondsElapsed++
				t.SetState(c)
			}
		}
	}()
}

func (t *TimerDef) ComponentWillUnmount() {
	t.State().ticker.Stop()
}

// Render renders the Timer component
func (t *TimerDef) Render() r.Element {
	return r.Div(nil,
		r.Div(nil,
			r.S(fmt.Sprintf("Seconds elapsed %.0f", t.State().secondsElapsed)),
		),
	)
}
