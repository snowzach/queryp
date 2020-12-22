package queryp

import "fmt"

type FilterOp int

const (
	FilterOpEquals           FilterOp = 0
	FilterOpNotEquals        FilterOp = 1
	FilterOpLessThan         FilterOp = 2
	FilterOpLessThanEqual    FilterOp = 3
	FilterOpGreaterThan      FilterOp = 4
	FilterOpGreaterThanEqual FilterOp = 5
	FilterOpLike             FilterOp = 6
	FilterOpNotLike          FilterOp = 7
	FilterOpILike            FilterOp = 8
	FilterOpNotILike         FilterOp = 9
	FilterOpRegexp           FilterOp = 10
	FilterOpNotRegexp        FilterOp = 11
	FilterOpIRegexp          FilterOp = 12
	FilterOpNotIRegexp       FilterOp = 13

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

	FieldNeedsQuote = `!=<>~:"`
)

var (
	FilterOpSymToFilterOp = map[string]FilterOp{
		FilterOpSymEquals:            FilterOpEquals,
		FilterOpSymNotEquals:         FilterOpNotEquals,
		FilterOpSymLessThan:          FilterOpLessThan,
		FilterOpSymGreaterThan:       FilterOpGreaterThan,
		FilterOpSymLessThanEqual:     FilterOpLessThanEqual,
		FilterOpSymLessThanEqual2:    FilterOpLessThanEqual,
		FilterOpSymGreaterThanEqual:  FilterOpGreaterThanEqual,
		FilterOpSymGreaterThanEqual2: FilterOpGreaterThanEqual,
		FilterOpSymLike:              FilterOpLike,
		FilterOpSymNotLike:           FilterOpNotLike,
		FilterOpSymILike:             FilterOpILike,
		FilterOpSymNotILike:          FilterOpNotILike,
		FilterOpSymRegexp:            FilterOpRegexp,
		FilterOpSymNotRegexp:         FilterOpNotRegexp,
		FilterOpSymIRegexp:           FilterOpIRegexp,
		FilterOpSymNotIRegexp:        FilterOpNotIRegexp,
	}
	FilterOpToFilterOpSym = map[FilterOp]string{
		FilterOpEquals:           FilterOpSymEquals,
		FilterOpNotEquals:        FilterOpSymNotEquals,
		FilterOpLessThan:         FilterOpSymLessThan,
		FilterOpGreaterThan:      FilterOpSymGreaterThan,
		FilterOpLessThanEqual:    FilterOpSymLessThanEqual,
		FilterOpGreaterThanEqual: FilterOpSymGreaterThanEqual,
		FilterOpLike:             FilterOpSymLike,
		FilterOpNotLike:          FilterOpSymNotLike,
		FilterOpILike:            FilterOpSymILike,
		FilterOpNotILike:         FilterOpSymNotILike,
		FilterOpRegexp:           FilterOpSymRegexp,
		FilterOpNotRegexp:        FilterOpSymNotRegexp,
		FilterOpIRegexp:          FilterOpSymIRegexp,
		FilterOpNotIRegexp:       FilterOpSymNotIRegexp,
	}
)

func (op FilterOp) MarshalJSON() ([]byte, error) {
	return []byte(`"` + op.String() + `"`), nil
}

func (op FilterOp) String() string {
	if filterOptSym, found := FilterOpToFilterOpSym[op]; found {
		return filterOptSym
	}
	return fmt.Sprintf(`OP(%d):`, op)
}
