package queryp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterAppend(t *testing.T) {

	var qp QueryParameters
	qp.Filter = append(qp.Filter, &FilterTerm{Logic: FilterLogicAnd, Field: "field1", Op: FilterOpEquals, Value: "1"})
	qp.Filter.Append(FilterLogicAnd, "field2", FilterOpNotEquals, "2").Append(FilterLogicOr, "field5", FilterOpLike, "value5%")
	qp.Options.Set("field3", "value3")
	qp.Options["field4"] = "value4"
	qp.Sort.Append("field2", false).Append("field3", true)
	qp.Sort = append(qp.Sort, &SortTerm{Field: "field4"})
	qp.Filter.SubFilter(FilterLogicAnd, NewFilter().Append(FilterLogicAnd, "field6", FilterOpLessThan, "6"))
	qp.Limit = 10
	qp.Offset = 20
	assert.Equal(t, `field1=1&field2!=2|field5=~value5%&(field6<6)&limit=10&offset=20&option[field3]=value3&option[field4]=value4&sort=field2,-field3,field4`, qp.String())

}
