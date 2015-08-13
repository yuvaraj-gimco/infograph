package tsdb

type Executor interface {
	Execute(queries QuerySlice, context *QueryContext) *BatchResult
}

func getExecutorFor(dsInfo *DataSourceInfo) Executor {
	switch dsInfo.Type {
	case "graphite":
		return &GraphiteExecutor{dsInfo}
	case "test":
		return unitTestExecutor
	}

	panic("No executor found for data source type: " + dsInfo.Type)
}

type UnitTestExecutor struct {
	results   map[string]*QueryResult
	resultsFn map[string]executorTestFunc
}

type executorTestFunc func(context *QueryContext) *QueryResult

var unitTestExecutor = &UnitTestExecutor{
	results: make(map[string]*QueryResult),
}

func (e *UnitTestExecutor) Execute(queries QuerySlice, context *QueryContext) *BatchResult {
	result := &BatchResult{QueryResults: make(map[string]*QueryResult)}
	for _, query := range queries {
		if results, has := e.results[query.RefId]; has {
			result.QueryResults[query.RefId] = results
		}
		if testFunc, has := e.resultsFn[query.RefId]; has {
			result.QueryResults[query.RefId] = testFunc(context)
		}
	}

	return result
}

type GraphiteExecutor struct {
	*DataSourceInfo
}

func (e *GraphiteExecutor) Execute(queries QuerySlice, context *QueryContext) *BatchResult {
	return &BatchResult{
		QueryResults: map[string]*QueryResult{},
	}
}
