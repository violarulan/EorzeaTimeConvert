# EorzeaTimeConvert

Convert Earth time to FFXIV Eorzea time

## Time

The measurement and expression of time in Eorzea, like the real world, is governed by a fundamental set of rules. The Dusty Tomes found in various guilds explain these standards in a book called The Five Ages - An Eorzean Chronology, albeit in heavily archaic language. In-game, time can be expressed in either a 24-hour clock format, or a 12-hour format with a.m./p.m. You can also change the clock to be Eorzea Time (ET), System Time (ST) of the server you're on, or Local Time (LT) showing your current real-life time. There is also an icon showing the current weather for the zone you're in, in the upper-right corner of the mini-map.

## Eorzean Time Units and Conversions

 Eorzean Increments | Components in Eorzean Units |  Conversion into Earth Time
--------------------|-----------------------------|----------------------------
       1 minute     |          60 seconds         |        2 11/12 seconds
        1 bell      |          60 minutes         |     2 minutes, 55 seconds
        1 sun       |           24 bells          |          70 minutes
        1 week      |            8 suns           |      9 hours, 20 minutes
        1 moon      |           32 suns           |      37 hours, 20 minutes
        1 year      |           12 moons          |        18 days, 16 hours 

## Usage

    package main

    import (
        "fmt"
        "time"
        c "github.com/violarulan/EorzeaTimeConvert"
    )

    func main(){
        // Convert a time.Time(Earth Time) to a EorzeaTime
        var ret = c.ConvertToEorzeaTime(time.Now())
        fmt.Println(ret)
        // {924 6 30 2 31 53}

        // Convert a time.Time(Earth Time) to a Eorzea time string 
        var str = c.ConvertToEorzeaTimeString(time.Now(), "%s-%s-%s %s:%s:%s")
        fmt.Println(str)
        // 924-06-30 08:01:22

        // Convert a Eorzea Time string to a time.Time(Earth Time)
        // For golang time format is restricted to "2006-01-02 15:04:05"
        // e.g: Convert 0924-07-15 11:49:01
        var time = c.ConvertToEarthTime("2006-01-02 15:04:05", "0924-07-15 11:49:01")
        fmt.Println(time)
        // 2017-03-14 17:54:27 +0900 JST
    }

## Contribute

Please feel free to submit issues and PRs.
