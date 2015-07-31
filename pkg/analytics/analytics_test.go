package analytics

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMetricQuery(t *testing.T) {

	Convey("When doing metric request", t, func() {
		request := &Request{
			Queries: map[string]*Query{
				{Query: "asd"},
			},
		}

		err := HandleRequest(request)
		So(err, ShouldBeNil)
	})
}
