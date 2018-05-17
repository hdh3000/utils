package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	flagBaseSpeedMPH                     = flag.Float64("speed", 2, "your speed, over flat ground in mph")
	flagDistanceMiles                    = flag.Float64("dist", 0, "the horizontal distance you plan to hike in miles")
	flagVerticalClimbFeet                = flag.Float64("ascent", 0, "the total vertical distance you expect to ascend")
	flagComfortableVerticalDescentFeet   = flag.Float64("descent", 0, "the total vertical distance you expect to descend (<12 deg slope)")
	flagUnComfortableVerticalDescentFeet = flag.Float64("treacherous-descent", 0, "the total >12 def slope vertical distance you expect to descend")
	flagErrorInBaseRate                  = flag.Float64("error", .1, "how much error to introduce into your base rate")
	flagTimeBetweenBreaks                = flag.Float64("time-between-breaks", 2, "how many hours between each break")
	flagWillStopForLunch                 = flag.Bool("lunch", true, "will you stop for lunch?")
)

const (
	feetInMile                               = float64(5280)
	minutesBackPer1000ftDescent              = 10
	minutesLostPer1000ftUncomfortableDescent = 10
	hourOfVerticalDistance                   = 2000
	minutesPerBreak                          = 10
)

func main() {
	flag.Parse()
	if *flagDistanceMiles == 0 {
		log.Fatalf("Must Include distance (use -h for more details)")
	}

	// Naismith's rule is `Allow one hour for every 3 miles (5 km) forward, plus an additional hour for every 2,000 feet (600 m) of ascent.`

	// This implies a time equivalency between 3 miles (or whatever rate you are traveling at) horizontal distance, and 2000ft of vertical distance
	// Some guy named scarf had the smart idea to look at it like this

	//	equivalent distance = x + α·y
	//  where:
	//	x = horizontal distance
	//	y = vertical distance
	//	α = 7.92 (3 mi / 2000 ft)

	alpha := (*flagBaseSpeedMPH * feetInMile) / hourOfVerticalDistance

	horizontalTime := *flagDistanceMiles / *flagBaseSpeedMPH * float64(60)
	verticalTime := (alpha * *flagVerticalClimbFeet / feetInMile) / *flagBaseSpeedMPH * float64(60)

	naismithTime := horizontalTime + verticalTime

	// Now the problem with naismithTime is that it doesn't account for time gained because of DESCENT, which can be substantial.
	// Some other smart guy named Aitken, came up with some adjustments for vertical descent
	// Comfortable descent (5-12 degree) gets you back 10 minutes per 1000 ft
	// Uncomfortable descent (>12 degree) costs you an extra 10 minutes per 1000ft

	descentTimeBack := (*flagComfortableVerticalDescentFeet / 1000) * minutesBackPer1000ftDescent
	descentTimeLost := (*flagUnComfortableVerticalDescentFeet / 1000) * minutesLostPer1000ftUncomfortableDescent

	actualTime := naismithTime + descentTimeLost - descentTimeBack

	// Now we add in breaks (rounding to whole breaks)
	numBreaks := int((actualTime / (*flagTimeBetweenBreaks * 60)) + .5)
	breakTimeMins := numBreaks * minutesPerBreak

	actualTime += float64(breakTimeMins)

	if *flagWillStopForLunch {
		actualTime += 45
	}

	uncertain := actualTime * *flagErrorInBaseRate

	fmt.Printf("\n")
	fmt.Printf("** Calculations **\n")
	fmt.Printf("Distance:             +%s (%.2f miles) \n", time.Duration(horizontalTime)*time.Minute, *flagDistanceMiles)
	fmt.Printf("Ascent:               +%s (%.0f ft)\n", time.Duration(verticalTime)*time.Minute, *flagVerticalClimbFeet)
	fmt.Printf("Descent:              -%s (%.0f ft)\n", time.Minute*time.Duration(descentTimeBack), *flagComfortableVerticalDescentFeet)
	fmt.Printf("Treacherous Descent:  +%s (%.0f ft)\n", time.Minute*time.Duration(descentTimeLost), *flagUnComfortableVerticalDescentFeet)
	fmt.Printf("Breaks:               +%s (%d breaks)\n", time.Minute*time.Duration(breakTimeMins), numBreaks)
	if *flagWillStopForLunch {
		fmt.Printf("Lunch                 +45m0s\n")
	}
	fmt.Printf("\n")

	fmt.Printf("** Hike time esitmate ** \n")
	fmt.Printf("Min:      %s\n", time.Minute*time.Duration(actualTime-uncertain))
	fmt.Printf("Expected: %s\n", time.Minute*time.Duration(actualTime))
	fmt.Printf("Max:      %s\n", time.Minute*time.Duration(actualTime+uncertain))
	fmt.Printf("\n")

}
