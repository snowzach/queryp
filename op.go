package queryp

const (
	FilterOpSymEquals      = "="
	FilterOpSymNotEquals   = "!="
	FilterOpSymLessThan    = "<"
	FilterOpSymGreaterThan = ">"

	FilterOpSymLessThanEqual     = "<="
	FilterOpSymLessThanEqual2    = "=<"
	FilterOpSymGreaterThanEqual  = ">="
	FilterOpSymGreaterThanEqual2 = "=>"

	FilterOpSymLike     = "=~"
	FilterOpSymNotLike  = "!=~"
	FilterOpSymILike    = "=~~"
	FilterOpSymNotILike = "!=~~"

	FilterOpSymRegexp     = ":"
	FilterOpSymNotRegexp  = "!:"
	FilterOpSymIRegexp    = ":~"
	FilterOpSymNotIRegexp = "!:~"
)

var (
	FilterOpSymToFilterOp = map[string]FilterOp{
		FilterOpSymEquals:            FilterOp_Equals,
		FilterOpSymNotEquals:         FilterOp_NotEquals,
		FilterOpSymLessThan:          FilterOp_LessThan,
		FilterOpSymGreaterThan:       FilterOp_GreaterThan,
		FilterOpSymLessThanEqual:     FilterOp_LessThanEqual,
		FilterOpSymLessThanEqual2:    FilterOp_LessThanEqual,
		FilterOpSymGreaterThanEqual:  FilterOp_GreaterThanEqual,
		FilterOpSymGreaterThanEqual2: FilterOp_GreaterThanEqual,
		FilterOpSymLike:              FilterOp_Like,
		FilterOpSymNotLike:           FilterOp_NotLike,
		FilterOpSymILike:             FilterOp_ILike,
		FilterOpSymNotILike:          FilterOp_NotILike,
		FilterOpSymRegexp:            FilterOp_Regexp,
		FilterOpSymNotRegexp:         FilterOp_NotRegexp,
		FilterOpSymIRegexp:           FilterOp_IRegexp,
		FilterOpSymNotIRegexp:        FilterOp_NotIRegexp,
	}
	FilterOpToFilterOpSym = map[FilterOp]string{
		FilterOp_Equals:           FilterOpSymEquals,
		FilterOp_NotEquals:        FilterOpSymNotEquals,
		FilterOp_LessThan:         FilterOpSymLessThan,
		FilterOp_GreaterThan:      FilterOpSymGreaterThan,
		FilterOp_LessThanEqual:    FilterOpSymLessThanEqual,
		FilterOp_GreaterThanEqual: FilterOpSymGreaterThanEqual,
		FilterOp_Like:             FilterOpSymLike,
		FilterOp_NotLike:          FilterOpSymNotLike,
		FilterOp_ILike:            FilterOpSymILike,
		FilterOp_NotILike:         FilterOpSymNotILike,
		FilterOp_Regexp:           FilterOpSymRegexp,
		FilterOp_NotRegexp:        FilterOpSymNotRegexp,
		FilterOp_IRegexp:          FilterOpSymIRegexp,
		FilterOp_NotIRegexp:       FilterOpSymNotIRegexp,
	}
)

func (op FilterOp) MarshalJSON() ([]byte, error) {
	return []byte(`"` + op.String() + `"`), nil
}

// func (op FilterOp) String() string {
// 	if filterOptSym, found := FilterOpToFilterOpSym[op]; found {
// 		return filterOptSym
// 	}
// 	return fmt.Sprintf(`OP(%d):`, op)
// }
