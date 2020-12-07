package queryp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterAppend(t *testing.T) {

	var qp QueryParameters
	qp.Filter = qp.Filter.Append(FilterLogicAnd, "field1", FilterOpEquals, 1)
	qp.Options = qp.Options.Set("field", "value")
	assert.Equal(t, `{"filter":[{"logic":"AND","op":"=","field":"field1","value":"1"}],"sort":null,"options":{"field":"value"},"limit":0,"offset":0}`+"\n", qp.String())

}
