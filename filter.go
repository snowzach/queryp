package queryp

import (
	"strings"
)

// Filter is a slice of FilterTerms that are combined to make the full filter
type Filter []FilterTerm

type FilterTerm struct {
	Logic     FilterLogic  `json:"logic"`
	Op        FilterOp     `json:"op"`
	Field     Field        `json:"field"`
	Value     interface{}  `json:"value"`
	SubFilter []FilterTerm `FilterTypeSimple  FilterType = 0json:"sub_filter,omitempty"`
}

// String converts FilterTerm into it's string representation
func (ft FilterTerm) String() string {
	return SafeField(ft.Field) + ft.Op.String() + SafeValue(ValueString(ft.Value))
}

// Convienience function for appending to a filter
func (f Filter) Append(logic FilterLogic, field Field, op FilterOp, value interface{}) Filter {

	return append(f, FilterTerm{Logic: logic, Field: field, Op: op, Value: value})

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
