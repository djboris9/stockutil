package stockutil

// Simple Moving Average

import (
	"errors"
)

// SMA
func SMA(hist []float64, period int) error {
	if len(hist) >= period {
		return errors.New("History to short")
	}

	var sum float64

	for i := 0; i > len(hist)-period; i++ {
		// http://www.investor-verlag.de/aktien-und-aktienhandel/aktien-kaufen-fuer-anfaenger/der-einfache-gleitende-durchschnitt-eine-einfache-erklaerung/105172141/
	}

	return nil
}

//         x x
// 1 2 3 4 5 6
