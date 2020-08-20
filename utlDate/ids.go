package utlDate

import (
	"fmt"
	"time"
)

type Did = int64 // Did = DateId
type Date = time.Time

const (
	MinuteSeconds = 60

	DaySeconds = 24 * 3600
)

func newDate() Date {
	d := time.Now()
	a, b := d.Zone()
	fmt.Println(a, b, d.Hour(), int(d.Weekday()))
	return d
}

func getHourOffset(d Date) time.Duration {
	return time.Duration(d.Hour())*time.Hour +
		time.Duration(d.Minute())*time.Minute +
		time.Duration(d.Second())*time.Second +
		time.Duration(d.Nanosecond())*time.Nanosecond

}

func GetUtcDayId(d Date) Did {
	d.Add(-getHourOffset(d))
	return d.Unix() / DaySeconds
}

func FromUtcDayId(dayId Did) Date {
	return time.Unix(dayId*DaySeconds, 0)
}

func getMonday(d Date) Date {
	off := int(d.Weekday()) - 1
	if off < 0 {
		off += 7
	}
	return d.AddDate(0, 0, -off)
}

func GetWeekId(d Date) Did {
	d = getMonday(d)
	day := GetUtcDayId(d)
	return (day + 3) / 7
}

func FromWeekId(weekId Did) Date {
	return FromUtcDayId(weekId*7 - 3)
}
