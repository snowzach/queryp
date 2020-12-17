package queryp

const (
	FilterLogicSymAnd = "&"
	FilterLogicSymOr  = "|"
)

var (
	FilterLogicSymToFilterLogic = map[string]FilterLogic{
		FilterLogicSymAnd: FilterLogic_And,
		FilterLogicSymOr:  FilterLogic_Or,
	}
)

func (logic FilterLogic) MarshalJSON() ([]byte, error) {
	return []byte(`"` + logic.String() + `"`), nil
}

// func (logic FilterLogic) String() string {
// 	switch logic {
// 	case FilterLogic_And:
// 		return "AND"
// 	case FilterLogic_Or:
// 		return "OR"
// 	default:
// 		return fmt.Sprintf(`LOGIC(%d)`, logic)
// 	}
// }
