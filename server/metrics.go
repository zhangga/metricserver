package server

const (
	Store = iota
	Counter
	Timer
	Rate
)

type MetricData struct {
	Type	int
	Name	string
	Tags	string
	Value	float64
	Time	int64
}
