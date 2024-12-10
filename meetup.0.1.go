/* My goal is to write readable code first. My second goal
is to write clean code (whatever that means.) I am happy to 
take feedback for those goals first. Efficiency is secondary
at this point.*/
package meetup
import (
    "fmt"
    "time"
    //"math"
	)
// Define the WeekSchedule type here.
type WeekSchedule int
type Weekday int
const (
    First WeekSchedule = 1
    Second WeekSchedule =  2
    Third WeekSchedule = 3
    Fourth WeekSchedule = 4
    Last WeekSchedule = -6
    Teenth WeekSchedule = 13
)
const (   
    Sunday			Weekday = iota
	Monday        	//Weekday = 1    
	Tuesday       	//Weekday = 2           
	Wednesday		//Weekday = 3
	Thursday		//Weekday = 4
	Friday			//Weekday = 5
	Saturday  		//Weekday = 6 
)
func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) (day int) {
	fmt.Printf("wSched: %v wDay: %v month %v year: %v\n", wSched, wDay, month, year)
    i:=0
    if wSched == Teenth {
        dateStr := fmt.Sprintf("%v-%d-%d", month, wSched, year)
        layout := "January-2-2006"
        parsedDate, err := time.Parse(layout, dateStr)
    if err != nil {
        fmt.Println("Error:", err)
        return 0
    	}
        parsedWeekday := Weekday(parsedDate.Weekday())
        for Weekday(wDay) != parsedWeekday && i < 10{
            parsedWeekday++
            if parsedWeekday > 6 {
                parsedWeekday = 0
            }
            i++
        }
        day = int(wSched) + i
    } else if wSched == Last {
        date := 1
        if month == 12 {
            month = 1
            year++
        } else {
            month++
        }
        dateStr := fmt.Sprintf("%v-%d-%d", month, date, year)
        layout := "January-2-2006"
        dayToCheck, err := time.Parse(layout, dateStr)
        if err != nil {
            fmt.Println("Error:", err)
            return 0
    	}
        dayToCheck = dayToCheck.AddDate(0,0,-7)
        date = int(dayToCheck.Day())
        fmt.Printf("dayToCheck %v\n", dayToCheck)
        i := 0
        parsedWeekday := Weekday(dayToCheck.Weekday())
        for Weekday(wDay) != parsedWeekday && i < 7 {
            parsedWeekday++
            if parsedWeekday > 6 {
                parsedWeekday = 0
            }
            i++
        }
        day = date + i
    } else {
        date := 1
        dateStr := fmt.Sprintf("%v-%d-%d", month, date, year)
        layout := "January-2-2006"
        dayToCheck, err := time.Parse(layout, dateStr)
        if err != nil {
            fmt.Println("Error:", err)
            return 0
        	}
        found := 0
        for found != int(wSched) && date < 31 {
            dateStr = fmt.Sprintf("%v-%d-%d", month, date, year)
            dayToCheck, _ = time.Parse(layout, dateStr)
            if dayToCheck.Weekday() == wDay {
                found++
            }
            date++
        }
        day = int(dayToCheck.Day())
    }
    return day
}