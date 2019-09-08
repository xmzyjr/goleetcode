package game_19_9_8

func dayOfTheWeek(day int, month int, year int) string {
	weeks := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	days := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		days[1] = 29
	} else {
		days[1] = 28
	}
	sum := 0
	for i := 0; i < month-1; i++ {
		sum += days[i]
	}
	sum += day
	return weeks[sum%7]
}
