package queryp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterAppend(t *testing.T) {

	var qp QueryParameters
	qp.Filter = append(qp.Filter, &FilterTerm{Logic: FilterLogic_And, Field: "field1", Op: FilterOp_Equals, Value: "1"})
	qp.SetOption("field", "value")
	assert.Equal(t, `{"filter":[{"logic":"AND","op":"=","field":"field1","value":"1"}],"sort":null,"options":{"field":"value"},"limit":0,"offset":0}`+"\n", qp.String())

}
