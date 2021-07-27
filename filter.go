package queryp

import (
	"strings"
)

const NullValue = "null"

// Filter is a slice of FilterTerms that are combined to make the full filter
type Filter []*FilterTerm

type FilterTerm struct {
	Logic     FilterLogic   `json:"logic"`
	Op        FilterOp      `json:"op"`
	Field     Field         `json:"field"`
	Value     interface{}   `json:"value"`
	SubFilter []*FilterTerm `json:"sub_filter,omitempty"`
}

// Creates a new filter
func NewFilter() *Filter {
	filter := make(Filter, 0)
	return &filter
}

// String converts FilterTerm into it's string representation
func (ft FilterTerm) String() string {
	return SafeField(ft.Field) + ft.Op.String() + ValueString(ft.Value)
}

// Convienience function for appending to a filter
func (f *Filter) Append(logic FilterLogic, field Field, op FilterOp, value interface{}) *Filter {
	if *f == nil {
		*f = make([]*FilterTerm, 0)
	}
	*f = append(*f, &FilterTerm{Logic: logic, Field: field, Op: op, Value: value})
	return f
}

// Convienience function for appending a sub filter
func (f *Filter) SubFilter(logic FilterLogic, subFilter *Filter) *Filter {
	if *f == nil {
		*f = make([]*FilterTerm, 0)
	}
	*f = append(*f, &FilterTerm{Logic: logic, SubFilter: *subFilter})
	return f
}

// Conveinience function for assigning to a qp.Filter
func (f *Filter) Filter() Filter {
	return *f
}

// Filter turns it into it's string representations
func (f Filter) String() string {

	var sb strings.Builder

	for i, ft := range f {
		if i > 0 {
			sb.WriteString(ft.Logic.String())
		}
		if len(ft.SubFilter) > 0 {
			_, _ = sb.WriteString("(")
			_, _ = sb.WriteString(Filter(ft.SubFilter).String())
			_, _ = sb.WriteString(")")
		} else {
			sb.WriteString(ft.String())
		}
	}

	return sb.String()

}
