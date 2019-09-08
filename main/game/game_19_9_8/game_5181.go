package game_19_9_8

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	var leftToRight, rightToLeft int
	for i := 0; i < len(distance); i++ {
		if i >= start && i < destination {
			leftToRight += distance[i]
		} else {
			rightToLeft += distance[i]
		}
	}
	if leftToRight < rightToLeft {
		return leftToRight
	} else {
		return rightToLeft
	}
}
