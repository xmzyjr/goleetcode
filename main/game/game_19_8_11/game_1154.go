package game_19_8_11

import (
	"strconv"
	"strings"
)

func ordinalOfDate(date string) int {
	if len(date) < 10 {
		return 0
	}
	//month := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	days1 := []int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	split := strings.Split(date, "-")
	i1, e := strconv.Atoi(split[0])
	i2, e2 := strconv.Atoi(split[1])
	i3, e1 := strconv.Atoi(split[2])
	if e != nil || e1 != nil || e2 != nil {
		return 0
	}
	var count int
	if i1%400 == 0 || (i1%4 == 0 && i1%100 != 0) {
		count = getNum(days1, i2, i3)
	} else {
		count = getNum(days, i2, i3)
	}
	return count

}

func getNum(day []int, i2 int, i3 int) int {
	count := 0
	for i := 1; i < i2; i++ {
		count += day[i]
	}
	count += i3
	return count
}
