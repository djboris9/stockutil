package stockutil

import (
	"os"
	"reflect"
	"testing"
)

// Test sets
var s1 = []float64{}
var s2 = []float64{0}
var s3 = []float64{4.2}
var s4 = []float64{1, 2}
var s5 = []float64{2, 1, 1.5}
var s6 = []float64{25, 85, 65, 45, 95, 75, 15, 35}
var s7 []float64 // Initialized in TestMain

func TestMain(m *testing.M) {
	s7 = make([]float64, len(testseries))
	for i, v := range testseries {
		// v[1] --> Price
		s7[i] = v[1]
	}

	os.Exit(m.Run())
}

func TestAVG(t *testing.T) {
	if AVG(s1) != 0 {
		t.Error("s1")
	}

	if AVG(s2) != 0 {
		t.Error("s2")
	}

	if AVG(s3) != 4.2 {
		t.Error("s3")
	}

	if AVG(s4) != 1.5 {
		t.Error("s4")
	}

	if AVG(s5) != 1.5 {
		t.Error("s5")
	}

	if AVG(s6) != 55 {
		t.Error(AVG(s6))
	}

	if AVG(s7) != 438.54194444444505 {
		t.Error("s7")
	}
}

func TestEMA(t *testing.T) {
	if _, err := EMA(s6, 9); err == nil {
		t.Error("Should be exception")
	}

	if v, _ := EMA(s6, 7); v[0] != 57.857142857142854 || len(v) != 2 {
		t.Error(v)
	}

	if v, _ := EMA(s7, 12); v[12] != 438.7392070304274 || v[80] != 438.81968204117805 || v[180] != 439.14635595306703 || len(v) != 889 {
		t.Errorf("%v %v %v %v\n", v[12], v[80], v[180], len(v))
	}
}

func TestSMA(t *testing.T) {
	if _, err := SMA(s1, 1); err == nil {
		t.Error("Should be exception")
	}

	if _, err := SMA(s1, 1); err == nil {
		t.Error("Should be exception")
	}

	if _, err := SMA(s3, 2); err == nil {
		t.Error("Should be exception")
	}

	if x, _ := SMA(s3, 1); x[0] != 4.2 || len(x) != 1 {
		t.Error(x)
	}

	// S6
	s6_res := []float64{58.333333333333336, 65, 68.33333333333333, 71.66666666666667, 61.666666666666664, 41.666666666666664}

	s6_try, _ := SMA(s6, 3)
	if !reflect.DeepEqual(s6_try, s6_res) {
		t.Errorf("%v should be %v\n", s6_try, s6_res)
	}
}

func TestMACD(t *testing.T) {
	// Test data stolen from http://investexcel.net/how-to-calculate-macd-in-excel/ , Sorry ;-)
	k := []float64{459.99, 448.85, 446.06, 450.81, 442.80, 448.97, 444.57, 441.40, 430.47, 420.05, 431.14, 425.66, 430.58, 431.72, 437.87, 428.43, 428.35, 432.50, 443.66, 455.72, 454.49, 452.08, 452.73, 461.91, 463.58, 461.14, 452.08, 442.66, 428.91, 429.79, 431.99, 427.72, 423.20, 426.21, 426.98, 435.69, 434.33, 429.80, 419.85, 426.24}

	macd, signal, err := MACD(k, 12, 26, 9)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	if len(macd) != 15 || len(signal) != 7 {
		t.Errorf("%v %v\n", len(macd), len(signal))
	}

	if macd[0] != 8.275269503907623 || macd[6] != 0.10298149119910249 {
		t.Errorf("%v %v\n", macd[0], macd[6])
	}

	if signal[0] != 3.037525868733945 || signal[6] != -1.338100412582993 {
		t.Errorf("%v %v\n", signal[0], signal[6])
	}
}
