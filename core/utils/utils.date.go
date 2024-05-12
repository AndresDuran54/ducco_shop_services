package utils

import (
	"fmt"
	"time"
)

type DateInformation struct {
	Original    time.Time
	Date        string
	DateTime    string
	DateTimeNew string
	TimeStamp   uint64
	Error       error
}

type UtilDate struct {
}

func (obj UtilDate) CurrentTimeUTC() DateInformation {
	now := obj.Now().UTC()

	return DateInformation{
		Error:     nil,
		Original:  now,
		Date:      obj.Format(now, "2006-01-02"),
		DateTime:  obj.FormatMilliseconds(now),
		TimeStamp: uint64(now.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))),
	}
}

func (obj UtilDate) Now() time.Time {
	return time.Now()
}

func (obj UtilDate) FormatMilliseconds(date time.Time) string {
	return obj.Format(date, "2006-01-02 15:04:05.000")
}

func (obj *UtilDate) Format(date time.Time, format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05.000"
	}

	var str string

	if !date.IsZero() {
		str = date.Format(format)
	} else {
		str = time.Now().Format(format)
	}

	return str
}

func (obj UtilDate) AddMinutes(date time.Time, value int) DateInformation {
	return obj.add(date, time.Duration(value)*time.Minute)
}

func (obj UtilDate) AddHours(date time.Time, value int) DateInformation {
	return obj.add(date, time.Duration(value)*time.Hour)
}

func (obj UtilDate) add(date time.Time, period time.Duration) DateInformation {
	finalDate := date

	if date.IsZero() {
		return DateInformation{Error: fmt.Errorf("error date is zero")}
	}

	finalDate = date.Add(period)

	return DateInformation{
		Error:     nil,
		Original:  date,
		Date:      obj.Format(finalDate, "2006-01-02"),
		DateTime:  obj.FormatMilliseconds(finalDate),
		TimeStamp: uint64(finalDate.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))),
	}
}
