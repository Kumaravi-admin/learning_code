/*
              Time & Computers

• Wall clock time represents real-world time.
• Wall clock time can jump forwards and  backwards.
• It gives you specific time relative to standard like UTC.


                Monotonic Time

• A monotonic time clocks is an ever increasing value.
• Used to measure elapsed time since.
• Monotonic time is used to get elapsed time while wall clock time gives the real-world time.


          Standard lib time package

• Go provides support for time with standard lib time package.
• The representation of time in Go uses both the wall clock and the monotonic clock.
• Two main type that Go provides to work with time are Time and Duration.


                   time.Duration


• Duration is represented as span of time with int64 nanosecond count.
• Duration constants defined in the time package:
      Nanosecond  Duration = 1
      Microsecond          = 1000 * Nanosecond
      Millisecond          = 1000 * Microsecond
      Second               = 1000 * Millisecond
      Minute               = 60 * Second
      Hour                 = 60 * Minute

              Example & Parsing


• Example of duration:
    ◦ d := 2 * time.Hour + 15 * time.Minute
• Duration can be parsed using the  ParseDuration function.
• Since time.Duration type is based on int64, you can use comparison operators with it.


             time.Duration Methods

• Duration also have methods to get  individual time components like hours, minutes, seconds, etc.
• Functions Since() & Until() take Time and  return Duration.


                   time.Time

• Time struct is used to represent a specific point in time.
• It stores year, month, day, hour, minute, seconds, nanoseconds along with the location  information.
• You should use Time as value and not as  pointers.
• The Sub() methods subtracts two instants of Time to produce Duration, while the Add() method adds Time and a Duration to produce Time.


              Zero Value & Comparison

• The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
• The IsZero() method should be used to check for the zero value of time.
• The equality operator (==) compares the time instant but also associated Location and monotonic clock reading.
• Since time has location information, Time should be compared using the Before(), After() & Equal() methods.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	// Example 1: Get the current time.
	// time.Now() returns a time.Time value (wall clock + monotonic reading).
	now := time.Now()
	fmt.Println(now)

	// Example 2: Add a duration to a time.
	// 2 * time.Hour is a time.Duration (a span of time, stored as nanoseconds).
	twoHours := 2 * time.Hour

	// Time.Add(Duration) returns a NEW Time (the original "now" is not changed).
	futureTime := now.Add(twoHours)
	fmt.Println("Time 2 hours from now: ", futureTime)

	// Example 3: Measure how long some code takes (elapsed time).
	// Step 1: Note the start time.
	start := time.Now()

	// Step 2: Do some work. Here: a simple loop that runs 10 million times.
	a := 1
	for i := 0; i < 10000000; i++ {
		a++
	}

	// Step 3: time.Since(start) = time.Now() minus start. It returns a Duration.
	// It uses the monotonic clock, so the result is correct even if the
	// wall clock changes in between (slide: "Monotonic Time").
	elapsed := time.Since(start)
	fmt.Println("Loop execution time: ", elapsed)
}

/*  Terminal output (your values will be different):

2026-11-06 11:13:19.9397383 +0530 IST m=+0.000524601
Time 2 hours from now:  2026-11-06 13:13:19.9397383 +0530 IST m=+7200.000524601
Loop execution time:  4.1797ms

*/