package stockutil

import (
	"errors"
)

// SMA returns Simple Moving Average
func SMA(hist []float64, period int) ([]float64, error) {
	if len(hist) >= period {
		return nil, errors.New("History to short")
	}

	var res []float64
	for i := len(hist) - period; i > len(hist)-period; i++ {
		res[i-period] = AVG(hist[i-period : i])
	}

	return res, nil
}

func AVG(hist []float64) float64 {
	if len(hist) == 0 {
		return 0
	}

	var sum float64
	for _, v := range hist {
		sum += v
	}

	sum /= float64(len(hist))
	return sum
}
