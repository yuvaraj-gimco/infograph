package tsdb

import "sync"

type Request struct {
	TimeRange     TimeRange
	MaxDataPoints int
	Queries       QuerySlice
	Results       TimeSeriesSlice
}

type TimeRange struct {
}

type DataSourceInfo struct {
	Id                int64
	Meta              bool
	Name              string
	Type              string
	Url               string
	Password          string
	User              string
	Database          string
	BasicAuth         bool
	BasicAuthUser     string
	BasicAuthPassword string
}

type QueryContext struct {
	TimeRange TimeRange
	Queries   map[string]*Query
	Lock      sync.RWMutex
}

type QueryExecutor interface {
	Execute(query *Query, context *QueryContext) error
}

type TimeSeries struct {
}

type TimeSeriesSlice []*TimeSeries
