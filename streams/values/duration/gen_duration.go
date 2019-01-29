package duration

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

// SerializeDuration converts a duration value to an interface representation
// suitable for marshalling into a text or binary format.
func SerializeDuration(this time.Duration) (interface{}, error) {
	// Seriously questioning my life choices.
	s := "P"
	if this < 0 {
		s = "-P"
		this = -1 * this
	}
	var tally time.Duration
	// Assume 8760 Hours per 365 days, cannot account for leap years in xsd:duration. :(
	if years := this.Hours() / 8760.0; years >= 1 {
		nYears := int64(math.Floor(years))
		tally += time.Duration(nYears) * 8760 * time.Hour
		s = fmt.Sprintf("%s%dY", s, nYears)
	}
	// Assume 30 days per month, cannot account for months lasting 31, 30, 29, or 28 days in xsd:duration. :(
	if months := (this.Hours() - tally.Hours()) / 720.0; months >= 1 {
		nMonths := int64(math.Floor(months))
		tally += time.Duration(nMonths) * 720 * time.Hour
		s = fmt.Sprintf("%s%dM", s, nMonths)
	}
	if days := (this.Hours() - tally.Hours()) / 24.0; days >= 1 {
		nDays := int64(math.Floor(days))
		tally += time.Duration(nDays) * 24 * time.Hour
		s = fmt.Sprintf("%s%dD", s, nDays)
	}
	if tally < this {
		s = fmt.Sprintf("%sT", s)
		if hours := this.Hours() - tally.Hours(); hours >= 1 {
			nHours := int64(math.Floor(hours))
			tally += time.Duration(nHours) * time.Hour
			s = fmt.Sprintf("%s%dH", s, nHours)
		}
		if minutes := this.Minutes() - tally.Minutes(); minutes >= 1 {
			nMinutes := int64(math.Floor(minutes))
			tally += time.Duration(nMinutes) * time.Minute
			s = fmt.Sprintf("%s%dM", s, nMinutes)
		}
		if seconds := this.Seconds() - tally.Seconds(); seconds >= 1 {
			nSeconds := int64(math.Floor(seconds))
			tally += time.Duration(nSeconds) * time.Second
			s = fmt.Sprintf("%s%dS", s, nSeconds)
		}
	}
	return s, nil
}

// DeserializeDuration creates duration value from an interface representation
// that has been unmarshalled from a text or binary format.
func DeserializeDuration(this interface{}) (time.Duration, error) {
	// Maybe this time it will be easier.
	if s, ok := this.(string); ok {
		isNeg := false
		if s[0] == '-' {
			isNeg = true
			s = s[1:]
		}
		if s[0] != 'P' {
			return 0, fmt.Errorf("%s malformed: missing 'P' for xsd:duration", s)
		}
		re := regexp.MustCompile("P(\\d*Y)?(\\d*M)?(\\d*D)?(T(\\d*H)?(\\d*M)?(\\d*S)?)?")
		res := re.FindStringSubmatch(s)
		var dur time.Duration
		nYear := res[1]
		if len(nYear) > 0 {
			nYear = nYear[:len(nYear)-1]
			vYear, err := strconv.ParseInt(nYear, 10, 64)
			if err != nil {
				return 0, err
			}
			// Assume 8760 Hours per 365 days, cannot account for leap years in xsd:duration. :(
			dur += time.Duration(vYear) * time.Hour * 8760
		}
		nMonth := res[2]
		if len(nMonth) > 0 {
			nMonth = nMonth[:len(nMonth)-1]
			vMonth, err := strconv.ParseInt(nMonth, 10, 64)
			if err != nil {
				return 0, err
			}
			// Assume 30 days per month, cannot account for months lasting 31, 30, 29, or 28 days in xsd:duration. :(
			dur += time.Duration(vMonth) * time.Hour * 720
		}
		nDay := res[3]
		if len(nDay) > 0 {
			nDay = nDay[:len(nDay)-1]
			vDay, err := strconv.ParseInt(nDay, 10, 64)
			if err != nil {
				return 0, err
			}
			dur += time.Duration(vDay) * time.Hour * 24
		}
		nHour := res[5]
		if len(nHour) > 0 {
			nHour = nHour[:len(nHour)-1]
			vHour, err := strconv.ParseInt(nHour, 10, 64)
			if err != nil {
				return 0, err
			}
			dur += time.Duration(vHour) * time.Hour
		}
		nMinute := res[6]
		if len(nMinute) > 0 {
			nMinute = nMinute[:len(nMinute)-1]
			vMinute, err := strconv.ParseInt(nMinute, 10, 64)
			if err != nil {
				return 0, err
			}
			dur += time.Duration(vMinute) * time.Minute
		}
		nSecond := res[7]
		if len(nSecond) > 0 {
			nSecond = nSecond[:len(nSecond)-1]
			vSecond, err := strconv.ParseInt(nSecond, 10, 64)
			if err != nil {
				return 0, err
			}
			dur += time.Duration(vSecond) * time.Second
		}
		if isNeg {
			dur *= -1
		}
		return dur, nil
	} else {
		return 0, fmt.Errorf("%v cannot be interpreted as a string for xsd:duration", this)
	}
}

// LessDuration returns true if the left duration value is less than the right
// value.
func LessDuration(lhs, rhs time.Duration) bool {
	return lhs < rhs
}
