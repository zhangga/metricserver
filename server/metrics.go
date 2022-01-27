package server

const (
	Store = iota
	Counter
	Timer
	Rate
)

// MetricData 从SDK接收到的数据结构
type MetricData struct {
	Type  int32   `json:"type"`
	Name  string  `json:"name"`
	Tags  string  `json:"tags"`
	Value float64 `json:"value"`
	Time  int64   `json:"time"`
}
