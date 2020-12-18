package queryp

import "fmt"

const (
	FilterLogicSymAnd = "&"
	FilterLogicSymOr  = "|"
)

var (
	FilterLogicSymToFilterLogic = map[string]FilterLogic{
		FilterLogicSymAnd: FilterLogicAnd,
		FilterLogicSymOr:  FilterLogicOr,
	}
)

func (logic FilterLogic) MarshalJSON() ([]byte, error) {
	return []byte(`"` + logic.String() + `"`), nil
}

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
