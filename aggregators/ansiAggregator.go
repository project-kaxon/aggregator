package aggregators

import "fmt"

type Aggregator interface {
	getName() string
	getNameLong() string
	getResult() string
}

func AnsiAggregator() {
	id := "ANSI"
	name := "American National Standards Institute"
	fmt.Println(id, name)
}
