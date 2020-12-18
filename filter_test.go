package queryp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterAppend(t *testing.T) {

	var qp QueryParameters
	qp.Filter = append(qp.Filter, FilterTerm{Logic: FilterLogicAnd, Field: "field1", Op: FilterOpEquals, Value: "1"})
	qp.Options = qp.Options.Set("field", "value")
	qp.Limit = 10
	qp.Offset = 20
	assert.Equal(t, `field1=1&limit=10&offset=20&option[field]=value`, qp.String())

}
