package tsdb

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMetricQuery(t *testing.T) {

	Convey("When batches groups for query", t, func() {

		Convey("Given 3 queries for 2 data sources", func() {
			request := &Request{
				Queries: QuerySlice{
					{RefId: "A", Query: "asd", DataSource: &DataSourceInfo{Id: 1}},
					{RefId: "B", Query: "asd", DataSource: &DataSourceInfo{Id: 1}},
					{RefId: "C", Query: "asd", DataSource: &DataSourceInfo{Id: 2}},
				},
			}

			batches, err := getBatches(request)
			So(err, ShouldBeNil)

			Convey("Should group into two batches", func() {
				So(len(batches), ShouldEqual, 2)
			})
		})

		Convey("Given query 2 depends on query 1", func() {
			request := &Request{
				Queries: QuerySlice{
					{RefId: "A", Query: "asd", DataSource: &DataSourceInfo{Id: 1}},
					{RefId: "B", Query: "asd", DataSource: &DataSourceInfo{Id: 2}},
					{RefId: "C", Query: "#A / #B", DataSource: &DataSourceInfo{Id: 3, Meta: true}},
				},
			}

			batches, err := getBatches(request)
			So(err, ShouldBeNil)

			Convey("Should return three batch groups", func() {
				So(len(batches), ShouldEqual, 3)
			})

			Convey("Group 3 should have group 1 and 2 as dependencies", func() {
				So(batches[2].Depends[0], ShouldEqual, &batches[0].Done)
				So(batches[2].Depends[1], ShouldEqual, &batches[1].Done)
			})

		})
	})
}
