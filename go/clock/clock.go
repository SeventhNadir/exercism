package clock

import (
	"fmt"
	"strconv"
)

//Clock type in minutes
type Clock struct {
	minutes int
}

//New constructs clock object.
func New(hours, minutes int) Clock {
	clock := new(Clock)
	clock.minutes += 60 * hours
	clock.minutes += minutes
	return clock.rationalise()
}

//String method for representing our clock in a friendly time format (ie "12:30")
func (clock Clock) String() string {
	clock.minutes = clock.minutes % 1440
	minutes := clock.minutes % 60
	hours := clock.minutes / 60
	strMinutes := strconv.Itoa(minutes)
	strHours := strconv.Itoa(hours)
	if minutes < 10 {
		strMinutes = "0" + strMinutes
	}
	if hours < 10 {
		strHours = "0" + strHours
	}
	return fmt.Sprintf("%s:%s", strHours, strMinutes)
}

//Add method for adding minutes to our clocks
func (clock Clock) Add(minutes int) Clock {
	clock.minutes += minutes
	return clock.rationalise()
}

//Subtract methods for subtracting minutes from our clocks
func (clock Clock) Subtract(minutes int) Clock {
	clock.minutes -= minutes
	return clock.rationalise()
}

func (clock Clock) rationalise() Clock {
	clock.minutes = clock.minutes % 1440
	if clock.minutes < 0 {
		clock.minutes += 1440
	}
	return clock
}
