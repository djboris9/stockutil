package stockutil

import "errors"

// SMA returns Simple Moving Average
func SMA(hist []float64, period int) ([]float64, error) {
	if len(hist) < period {
		return nil, errors.New("History too short")
	}

	res := make([]float64, len(hist)-period+1)
	for i := period; i <= len(hist); i++ {
		res[i-period] = AVG(hist[i-period : i])
	}

	return res, nil
}

// AVG returns a simple average
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

// EMA calculates the Exponential Moving Average.
// Multiplyer same as needed for MACD (New Style).
func EMA(hist []float64, period int) ([]float64, error) {
	if len(hist) < period {
		return nil, errors.New("History too short")
	}

	var initHist = hist[0:period]
	var emaHist = hist[period:len(hist)]

	var previousEMA = AVG(initHist)

	result := make([]float64, len(emaHist)+1)
	result[0] = previousEMA

	for k, price := range emaHist {
		var multiplier float64 = 2 / float64(period+1)
		previousEMA = price*multiplier + previousEMA*(1-multiplier)
		result[k+1] = previousEMA
	}

	return result, nil
}

// MACD returns the Moving Average Convergence Divergence
// Returns MACD(t), SIGNAL(t), error
func MACD(hist []float64, fast, slow, signal int) ([]float64, []float64, error) {
	if fast >= slow {
		return nil, nil, errors.New("Fast cannot be >= slow")
	}

	// Calculate EMAs for MACD(t)
	emaFast, err := EMA(hist, fast)
	if err != nil {
		return nil, nil, err
	}

	emaSlow, err := EMA(hist, slow)
	if err != nil {
		return nil, nil, err
	}

	// Calculate result of MACD(t)
	for i, _ := range emaSlow {
		emaSlow[i] = emaFast[i+(slow-fast)] - emaSlow[i]
	}
	macd := emaSlow

	// Calculate signal
	sign, err := EMA(macd, signal)
	if err != nil {
		return nil, nil, err
	}

	return macd, sign, nil
}
