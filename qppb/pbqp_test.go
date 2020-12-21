package qppb

import (
	"testing"

	"github.com/snowzach/queryp"
	"github.com/stretchr/testify/assert"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func TestPBtoQP(t *testing.T) {

	value, _ := structpb.NewValue("value")

	pb := &QueryParameters{
		Filter: []*FilterTerm{
			{
				Logic: FilterLogic_FilterLogicAnd,
				Field: "test",
				Op:    FilterOp_FilterOpEquals,
				Value: value,
			},
		},
		Options: map[string]string{
			"option1": "value1",
		},
		Limit:  10,
		Offset: 20,
	}

	qp := &queryp.QueryParameters{
		Filter: queryp.Filter{
			{
				Logic: queryp.FilterLogicAnd,
				Field: "test",
				Op:    queryp.FilterOpEquals,
				Value: "value",
			},
		},
		Options: map[string]string{
			"option1": "value1",
		},
		Limit:  10,
		Offset: 20,
	}

	c := pb.QueryParameters()
	assert.Equal(t, qp, c)

}
