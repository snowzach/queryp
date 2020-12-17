package qppg

import (
	"strings"
	"testing"

	"github.com/snowzach/queryp"
	"github.com/stretchr/testify/assert"
)

func TestQueryParser(t *testing.T) {

	q, err := queryp.ParseQuery("field=value&((another=<value|yet=another1|limit=weee))|third=value&limit=10&option=beans&sort=test3,-another")
	assert.Nil(t, err)
	assert.Equal(t, int64(10), q.Limit)
	assert.Equal(t, int64(0), q.Offset)
	assert.Equal(t, []queryp.SortTerm{{Field: "test3", Desc: false}, {Field: "another", Desc: true}}, q.Sort)

	var queryClause strings.Builder
	var queryParams = []interface{}{}

	filterFields := queryp.FilterFieldTypes{
		"another":           queryp.FilterType_String,
		"yet":               queryp.FilterType_String,
		"field":             queryp.FilterType_String,
		"limit":             queryp.FilterType_Numeric,
		"third":             queryp.FilterFieldCustom{FieldName: "whoathird", FilterType: queryp.FilterType_String},
		"thing.id":          queryp.FilterType_String,
		"thing.name":        queryp.FilterType_String,
		"thing.description": queryp.FilterType_String,
	}

	sortFields := queryp.SortFields{
		"test1",
		"test2",
		"test3:test7",
		"another",
	}

	err = FilterQuery(filterFields, q.Filter, &queryClause, &queryParams)
	assert.Nil(t, err)

	err = SortQuery(sortFields, q.Sort, &queryClause, &queryParams)
	assert.Nil(t, err)

	assert.Equal(t, "field = $1 AND ((another <= $2 OR yet = $3 OR limit = $4)) OR whoathird = $5 ORDER BY test7, another DESC", queryClause.String())
	assert.Equal(t, []interface{}([]interface{}{"value", "value", "another1", "weee", "value"}), queryParams)

}
