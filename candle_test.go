package stockutil

import (
	"testing"
	"time"
)

func TestFlatCandleAppend(t *testing.T) {
	// TODO: Implement
	t.Skip()
}

func loadFlatCandle() *FlatCandle {
	var ticks []Tick
	for _, v := range testseries {
		t := Tick{
			Time:  time.Unix(int64(v[0]), 0),
			Value: v[1],
		}
		ticks = append(ticks, t)
	}
	dur, _ := time.ParseDuration("15m")
	return TicksToFlatCandle(ticks, dur)
}

func TestFlatCandleGet(t *testing.T) {
	fcandle := loadFlatCandle()
	c, _ := fcandle.Get(1)

	if c.Open != 439.47 || c.High != 439.49 || c.Low != 437.55 || c.Close != 438.93 ||
		c.Time != fcandle.Time[1] || c.Duration != fcandle.Duration {

		t.Errorf("%v\n", c)
	}
}

func TestTicksToFlatCandle(t *testing.T) {
	fcandle := loadFlatCandle()

	if x := len(fcandle.Open); x != 75 {
		t.Errorf("%v\n", x)
	}

	if x := fcandle.Open[1]; x != 439.47 {
		t.Errorf("%v\n", x)
	}
	if x := fcandle.High[1]; x != 439.49 {
		t.Errorf("%v\n", x)
	}
	if x := fcandle.Low[1]; x != 437.55 {
		t.Errorf("%v\n", x)
	}
	if x := fcandle.Close[1]; x != 438.93 {
		t.Errorf("%v\n", x)
	}

	tt, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2015-12-22 01:40:00 +0100 CET")
	if x := fcandle.Time[1]; x != tt {
		t.Errorf("%v\n", x)
	}

	dur, _ := time.ParseDuration("15m")
	if x := fcandle.Duration; x != dur {
		t.Errorf("%v\n", x)
	}
}
