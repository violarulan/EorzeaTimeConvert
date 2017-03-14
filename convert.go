package EorzeaTimeConvert

import (
        "fmt"
        "math"
        "strconv"
        "time"
)

const (
        YEAR   = 33177600
        MONTH  = 2764800
        DAY    = 86400
        HOUR   = 3600
        MINUTE = 60
        SECOND = 1
)

type EorzeaTime struct {
        YearVal   int64
        MonthVal  int64
        DayVal    int64
        HourVal   int64
        MinuteVal int64
        SecondVal int64
}

var (
        EORZEA_TIME_CONSTANT float64 = float64(3600) / float64(175)
)

func ConvertToEorzeaTime(time time.Time) EorzeaTime {
        earthTime := float64(time.Unix())
        eorzeaTime := int64(math.Floor(earthTime * EORZEA_TIME_CONSTANT))

        var ret EorzeaTime
        ret.YearVal = int64(math.Floor(float64(eorzeaTime / YEAR))) + 1
        ret.MonthVal = int64(math.Floor(float64(eorzeaTime / MONTH % 12))) + 1
        ret.DayVal = int64(math.Floor(float64(eorzeaTime / DAY % 32))) + 1
        ret.HourVal = int64(math.Floor(float64(eorzeaTime / HOUR % 24)))
        ret.MinuteVal = int64(math.Floor(float64(eorzeaTime / MINUTE % 60)))
        ret.SecondVal = int64(math.Floor(float64(eorzeaTime / SECOND % 60)))

        return ret
}

func ConvertToEorzeaTimeString(time time.Time, format string) string {

        earthTime := float64(time.Unix())
        eorzeaTime := int64(math.Floor(earthTime * EORZEA_TIME_CONSTANT))

        yearVal := "" + strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / YEAR))) + 1, 10)
        monthVal := formatZero(strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / MONTH % 12))) + 1, 10))
        dayVal := formatZero(strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / DAY % 32 ))) + 1, 10))
        hourVal := formatZero(strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / HOUR % 24))), 10))
        minuteVal := formatZero(strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / MINUTE % 60))), 10))
        secondVal := formatZero(strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / SECOND % 60))), 10))

        var ret string
        ret = fmt.Sprintf(format, yearVal, monthVal, dayVal, hourVal, minuteVal, secondVal)
        return ret
}

func parseEorzeaTimeString(timestring string, format string) (time.Time, error) {

        date, err := time.Parse(format, timestring)
        var ret time.Time
        if err != nil {
                return ret, err
        }
        return date, nil
}

func ConvertToEarthTime(timestring string, format string) (time.Time, error) {
        
        date, err := parseEorzeaTimeString(format, timestring)
        var ret time.Time
        var utc int64

        if err != nil {
                return ret, err
        }
        years   := int64(date.Year())
        months  := int64(date.Month())
        days    := int64(date.Day())
        hours   := int64(date.Hour())
        minutes := int64(date.Minute())
        seconds := int64(date.Second())

        utc = int64(float64((years - 1) * YEAR + (months - 1) * MONTH + (days - 1) * DAY + hours * HOUR + minutes * MINUTE + seconds ) / EORZEA_TIME_CONSTANT)
        ret = time.Unix(utc, 0)
        return ret, nil
}


func formatZero(str string) string {
        if len(str) == 1 {
                return "0" + str
        } else {
                return str
        }
}

