# Enable this for SVG output
# set output "test.svg"
# set term svg enhanced size 1200,800

set multiplot layout 2, 1 title "BTC/USD 22.12.2015 02:50-05:22"
set autoscale
set lmargin 12

set ylabel "Value"
set datafile missing '0'
set timefmt "%s"

set xdata time
set timefmt "%s"

# Don't plot so many points
set xrange [1450749000:1450758120]

unset xtics

# ticks[i], sma[i], ema[i], macd[i], signal[i]
plot "gnuplot.dat" using 1:2 with lines title "Ticks" lt rgb "black", \
     "gnuplot.dat" using 1:3 with lines title "SMA"   lt rgb "red", \
     "gnuplot.dat" using 1:4 with lines title "EMA"   lt rgb "blue"

set tmargin 0
set format x "%H:%M:%S"
set xtics
set xtics rotate

# ticks[i], sma[i], ema[i], macd[i], signal[i]
plot "gnuplot.dat" using 1:5 with lines title "MACD", \
     "gnuplot.dat" using 1:6 with lines title "SIGNAL"

unset multiplot
