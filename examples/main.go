package main

// Run with `go run data.go main.go`

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	su "github.com/djboris9/stockutil"
)

func main() {
	// Convert my testseries to slice of ticks
	// testseries: []{UnixTimestamp, TickValue}
	ticks := make([]float64, len(testseries))
	for i, _ := range testseries {
		ticks[i] = testseries[i][1]
	}

	// Now calculate all indicators, using "ticks"
	sma, _ := su.SMA(ticks, 3)
	ema, _ := su.EMA(ticks, 7)
	macd, signal, _ := su.MACD(ticks, 12, 26, 9)

	// Fix alignments
	sma = append(make([]float64, len(ticks)-len(sma)), sma...)
	ema = append(make([]float64, len(ticks)-len(ema)), ema...)
	macd = append(make([]float64, len(ticks)-len(macd)), macd...)
	signal = append(make([]float64, len(ticks)-len(signal)), signal...)

	// Print all values to "gnuplot.dat"
	fd, err := os.Create("gnuplot.dat")
	if err != nil {
		log.Fatal(err.Error())
	}
	for i, _ := range ticks {
		fd.Write([]byte(fmt.Sprintf("%v %v %v %v %v %v\n", testseries[i][0], ticks[i], sma[i], ema[i], macd[i], signal[i])))
	}
	fd.Close()

	// Start Gnuplot
	out, err := exec.Command("gnuplot", "-p", "script.gnuplot").CombinedOutput()
	fmt.Println(err)
	fmt.Println(string(out))
}
