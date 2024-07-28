package main

import "fmt"

func main() {
	WeekDay := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	WorkingDays := make([]string, 6)
	copy(WorkingDays, WeekDay[0:5])
	fmt.Println(WorkingDays)
	WeekDay = WeekDay[5:]
	fmt.Println(WeekDay)

}
