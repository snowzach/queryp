package queryp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterAppend(t *testing.T) {

	qp := NewQueryParameters()
	qp.Filter = qp.Filter.Append(FilterLogicStart, "field1", FilterOpEquals, 1)
	assert.Equal(t, `{"filter":[{"logic":"START","op":"=","field":"field1","value":"1"}],"sort":[],"options":[],"limit":0,"offset":0}`+"\n", qp.String())

}
