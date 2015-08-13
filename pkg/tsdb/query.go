package tsdb

type Query struct {
	RefId      string
	Query      string
	DataSource *DataSourceInfo
	Results    []*TimeSeries
	Exclude    bool
}

type QuerySlice []*Query
