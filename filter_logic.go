package queryp

import "fmt"

type FilterLogic int

const (
	FilterLogicAnd FilterLogic = 0
	FilterLogicOr  FilterLogic = 1

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
