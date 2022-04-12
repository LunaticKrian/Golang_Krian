package main

import "fmt"

func main() {
	//假设还有100天放假，还有几个星期？
	var days int = 100
	var week int = days / 7
	var day int = days % 7

	fmt.Printf("%d天后放假 === %d周%d天后放假", days, week, day)
}
