package clock

import "strconv"

const testVersion = 4

type Clock struct {
	h int
	m int
}

func New(hour, minute int) Clock {
	hour, minute = normalizeMinute(hour, minute)
	return Clock{
		h: hour,
		m: minute,
	}
}
func normalizeHour(hour int) int {
	if hour > 23 {
		hour = hour % 24
	}
	if hour < 0 {
		hour = 24 + hour
		hour = normalizeHour(hour)
	}
	return hour
}
func normalizeMinute(hour, minute int) (int, int) {
	if minute >= 60 {
		hour = hour + (minute / 60)
		minute = minute % 60
	}
	if minute < 0 {
		minute = 60 + minute
		hour--
		hour, minute = normalizeMinute(hour, minute)
	}
	hour = normalizeHour(hour)
	return hour, minute
}
func (c Clock) String() string {
	hour := strconv.Itoa(c.h)
	min := strconv.Itoa(c.m)
	if len(hour) < 2 {
		hour = "0" + hour
	}
	if len(min) < 2 {
		min = "0" + min
	}
	return hour + ":" + min
}

func (c Clock) Add(minutes int) Clock {
	c.h, c.m = normalizeMinute(c.h, c.m+minutes)
	return c
}
