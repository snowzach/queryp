package queryp

import "fmt"

const (
	FilterLogicSymAnd = "&"
	FilterLogicSymOr  = "|"

	ValueNeedsQuote = `&|)"`
)

var (
	FilterLogicSymToFilterLogic = map[string]FilterLogic{
		FilterLogicSymAnd: FilterLogicAnd,
		FilterLogicSymOr:  FilterLogicOr,
	}
)

func (logic FilterLogic) String() string {
	switch logic {
	case FilterLogicAnd:
		return FilterLogicSymAnd
	case FilterLogicOr:
		return FilterLogicSymOr
	default:
		return fmt.Sprintf(`LOGIC(%d)`, logic)
	}
}
