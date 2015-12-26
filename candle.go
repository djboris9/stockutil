package stockutil

import (
	"errors"
	"time"
)

type Tick struct {
	Time  time.Time
	Value float64
}

type Candle struct {
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Time     time.Time
	Duration time.Duration
}

type FlatCandle struct {
	Open     []float64
	High     []float64
	Low      []float64
	Close    []float64
	Time     []time.Time
	Duration time.Duration
}

func (c *FlatCandle) Get(i int) (Candle, error) {
	if len(c.Open) <= i {
		return Candle{}, errors.New("Out of bounds")
	}
	return Candle{
		Open:     c.Open[i],
		High:     c.High[i],
		Low:      c.Low[i],
		Close:    c.Close[i],
		Time:     c.Time[i],
		Duration: c.Duration,
	}, nil
}

func TicksToFlatCandle(ticks []Tick, duration time.Duration) *FlatCandle {
	r := new(FlatCandle)
	r.Duration = duration

	var curOpen float64
	var curHigh float64
	var curLow float64
	var curClose float64
	var curTime time.Time
	var newShift time.Time

	for i, _ := range ticks {
		var wasZero bool = false
		if newShift.IsZero() {
			newShift = ticks[i].Time
			wasZero = true
		}

		// New Candle range began
		if !newShift.After(ticks[i].Time) {
			if !wasZero {
				// Add old candle to FlatCandle
				r.Open = append(r.Open, curOpen)
				r.High = append(r.High, curHigh)
				r.Low = append(r.Low, curLow)
				r.Close = append(r.Close, curClose)
				r.Time = append(r.Time, curTime)
			}

			// Initialize new candle
			newShift = ticks[i].Time.Add(duration)
			curTime = ticks[i].Time
			curOpen = ticks[i].Value
			curHigh = ticks[i].Value
			curLow = ticks[i].Value
		}

		if curHigh < ticks[i].Value {
			curHigh = ticks[i].Value
		}
		if curLow > ticks[i].Value {
			curLow = ticks[i].Value
		}
		curClose = ticks[i].Value

	}
	return r
}

/*
 Is this realy needed?
 - Problem 1: Appending
// Returns Open, High, Low, Close as slices from an candle slice
func FlattenCandle(c []Candle) *FlatCandle {
	r := &FlatCandle{}
	r.Open = make([]float64, len(c))
	r.High = make([]float64, len(c))
	r.Low = make([]float64, len(c))
	r.Close = make([]float64, len(c))
	r.Time = make([]time.Time, len(c))

	if len(c) > 0 {
		r.Duration = c[0].Duration
	}

	for i, _ := range c {
		r.Open[i] = c[i].Open
		r.Low[i] = c[i].Low
		r.High[i] = c[i].High
		r.Close[i] = c[i].Close
		r.Time[i] = c[i].Time
	}

	return r
}

func UnflatCandle(c *FlatCandle) []Candle {
	r := make([]Candle, len(c.Open))

	for i, _ := range c.Open {
		r[i] = Candle{
			Open:     c.Open[i],
			High:     c.High[i],
			Low:      c.Low[i],
			Close:    c.Close[i],
			Time:     c.Time[i],
			Duration: c.Duration,
		}
	}

	return r
}
*/
