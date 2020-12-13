package qppg

import (
	"strings"
	"testing"

	"github.com/snowzach/queryp"
	"github.com/stretchr/testify/assert"
)

func TestQueryParser(t *testing.T) {

	q, err := queryp.ParseQuery("field=value&((another=<value|yet=another1|limit=weee))|third=value&limit=10&option=beans&sort=test,-another")
	assert.Nil(t, err)
	assert.Equal(t, int64(10), q.Limit)
	assert.Equal(t, int64(0), q.Offset)
	assert.Equal(t, []queryp.SortTerm{{Field: "test", Desc: false}, {Field: "another", Desc: true}}, q.Sort)

	var queryClause strings.Builder
	var queryParams = []interface{}{}

	filterFields := queryp.FilterFieldTypes{
		"another":           queryp.FilterTypeString,
		"yet":               queryp.FilterTypeString,
		"field":             queryp.FilterTypeString,
		"limit":             queryp.FilterTypeNumeric,
		"third":             queryp.FilterFieldCustom{FieldName: "whoathird", FilterType: queryp.FilterTypeString},
		"thing.id":          queryp.FilterTypeString,
		"thing.name":        queryp.FilterTypeString,
		"thing.description": queryp.FilterTypeString,
	}

	err = FilterQuery(filterFields, q.Filter, &queryClause, &queryParams)
	assert.Nil(t, err)
	assert.Equal(t, "field = $1 AND ((another <= $2 OR yet = $3 OR limit = $4)) OR whoathird = $5", queryClause.String())
	assert.Equal(t, []interface{}([]interface{}{"value", "value", "another1", "weee", "value"}), queryParams)

}
