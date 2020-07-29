package main

import "fmt"

func displayDayOfWeek(str string, _idx int) string {
	switch str {
	case "sat":
		// 次の条件を実行する
		// satの場合はHoridayと表示される
		fallthrough
	case "sun":
		return "Horiday"
	case "mon":
		return "Monday!"
	default:
		return "Weekday"
	}
}

func main() {
	dayOfWeeks := [4]string{"sat", "sun", "mon", "tue"}

	for i, d := range dayOfWeeks {
		fmt.Println(displayDayOfWeek(d, i))
	}
}
