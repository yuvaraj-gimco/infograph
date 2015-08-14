package dtos

type Query struct {
	RefId   bool       `json:"refId"`
	Query   string     `json:"query"`
	Depends [][]string `json:"depends"`
	DsId    int64      `json:"dsId"`
	Exclude bool       `json:"exclude"`
}

type MetricRequest struct {
	Queries []Query `json:"queries"`
}

type QueryResult struct {
	Error  error      `json:"error"`
	RefId  string     `json:"refId"`
	Series TimeSeries `json:"series"`
}

type TimeSeries struct {
	Name   string       `json:"name"`
	Points [][2]float64 `json:"points"`
}

type MetricResponse struct {
	Results []QueryResult `json:"queries"`
}
