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
        ret.yearVal = int64(math.Floor(float64(eorzeaTime / YEAR))) + 1
        ret.monthVal = int64(math.Floor(float64(eorzeaTime / MONTH % 12))) + 1
        ret.dayVal = int64(math.Floor(float64(eorzeaTime / DAY % 32))) + 1
        ret.hourVal = int64(math.Floor(float64(eorzeaTime / HOUR % 24)))
        ret.minuteVal = int64(math.Floor(float64(eorzeaTime / MINUTE % 60)))
        ret.secondVal = int64(math.Floor(float64(eorzeaTime / SECOND % 60)))

        return ret
}

func ConvertToEorzeaTimeString(time time.Time, format string) string {

        earthTime := float64(time.Unix())
        eorzeaTime := int64(math.Floor(earthTime * EORZEA_TIME_CONSTANT))

        yearVal := "" + strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / YEAR))) + 1, 10)
        monthVal := formatZero(strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / MONTH%12))) + 1, 10))
        dayVal := formatZero(strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / DAY % 32 ))) + 1, 10))
        hourVal := formatZero(strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / HOUR % 24))), 10))
        minuteVal := formatZero(strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / MINUTE % 60))), 10))
        secondVal := formatZero(strconv.FormatInt(int64(math.Floor(float64(eorzeaTime / SECOND % 60))), 10))

        var ret string
        ret = fmt.Sprintf(format, yearVal, monthVal, dayVal, hourVal, minuteVal, secondVal)
        return ret
}

func formatZero(str string) string {
        if len(str) == 1 {
                return "0" + str
        } else {
                return str
        }
}

