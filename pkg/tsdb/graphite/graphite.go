package graphite

import (
	"net/http"
	"net/url"
	"time"

	"github.com/grafana/grafana/pkg/tsdb"
)

type GraphiteExecutor struct {
	*tsdb.DataSourceInfo
}

func NewGraphiteExecutor(dsInfo *tsdb.DataSourceInfo) tsdb.Executor {
	return &GraphiteExecutor{dsInfo}
}

func init() {
	tsdb.RegisterExecutor("graphite", NewGraphiteExecutor)
}

func (e *GraphiteExecutor) Execute(queries tsdb.QuerySlice, context *tsdb.QueryContext) *tsdb.BatchResult {
	result := &tsdb.BatchResult{}

	params := url.Values{
		"from":          []string{"now-1h"},
		"until":         []string{"now"},
		"maxDataPoints": []string{"500"},
	}

	client := http.Client{Timeout: time.Duration(10 * time.Second)}
	_, err := client.PostForm("/render?", params)
	if err != nil {
		result.Error = err
		return result
	}

	return &tsdb.BatchResult{
		QueryResults: map[string]*tsdb.QueryResult{},
	}
}
