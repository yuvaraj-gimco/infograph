package api

import (
	"math/rand"
	"strconv"

	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/tsdb"
	_ "github.com/grafana/grafana/pkg/tsdb/graphite"
)

func GetTestMetrics(c *middleware.Context) {
	from := c.QueryInt64("from")
	to := c.QueryInt64("to")
	maxDataPoints := c.QueryInt64("maxDataPoints")
	stepInSeconds := (to - from) / maxDataPoints

	result := dtos.MetricQueryResultDto{}
	result.Data = make([]dtos.MetricQueryResultDataDto, 1)

	for seriesIndex := range result.Data {
		points := make([][2]float64, maxDataPoints)
		walker := rand.Float64() * 100
		time := from

		for i := range points {
			points[i][0] = walker
			points[i][1] = float64(time)
			walker += rand.Float64() - 0.5
			time += stepInSeconds
		}

		result.Data[seriesIndex].Target = "test-series-" + strconv.Itoa(seriesIndex)
		result.Data[seriesIndex].DataPoints = points
	}

	c.JSON(200, &result)
}

func GetMetricsRequest(c *middleware.Context, req dtos.MetricRequest) Response {
	tsdbReq := &tsdb.Request{
		Queries: make(tsdb.QuerySlice, len(req.Queries)),
	}

	for _, query := range req.Queries {
		tsdbReq.Queries[0] = &tsdb.Query{
			RefId:   query.RefId,
			Query:   query.Query,
			Depends: query.Depends,
			DataSource: &tsdb.DataSourceInfo{
				Id:   1,
				Url:  "http://localhost:8080",
				Type: "graphite",
			},
		}
	}

	rsp, err := tsdb.HandleRequest(tsdbReq)
	if err != nil {
		return ApiError(500, "Failed to process metric request", err)
	}

	// translate to api dto
	metricRes := dtos.MetricResponse{Results: make(map[string]*dtos.QueryResult)}

	for refId, res := range rsp.Results {
		queryRes := &dtos.QueryResult{
			RefId:  refId,
			Series: make([]dtos.TimeSeries, len(res.Series)),
		}

		if res.Error != nil {
			queryRes.Error = res.Error.Error()
		}

		for i, series := range res.Series {
			queryRes.Series[i].Name = series.Name
			queryRes.Series[i].Points = series.Points
		}
		metricRes.Results[refId] = queryRes
	}

	return Json(200, metricRes)
}
