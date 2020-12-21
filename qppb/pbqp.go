package qppb

import "github.com/snowzach/queryp"

func (qp *QueryParameters) QueryParameters() *queryp.QueryParameters {
	return &queryp.QueryParameters{
		Filter:  FilterQP(qp.Filter),
		Sort:    SortQP(qp.Sort),
		Options: qp.Options,
		Limit:   int64(qp.Limit),
		Offset:  int64(qp.Offset),
	}

}

func FilterQP(ftpb []*FilterTerm) queryp.Filter {
	if len(ftpb) == 0 {
		return nil
	}
	f := make(queryp.Filter, 0, len(ftpb))
	for _, ft := range ftpb {
		f = append(f, queryp.FilterTerm{
			Logic:     ft.Logic.FilterLogic(),
			Field:     ft.Field,
			Op:        ft.Op.FilterOp(),
			Value:     ft.Value.AsInterface(),
			SubFilter: FilterQP(ft.SubFilter),
		})
	}
	return f
}

func (fl FilterLogic) FilterLogic() queryp.FilterLogic {
	switch fl {
	case FilterLogic_FilterLogicAnd:
		return queryp.FilterLogicAnd
	case FilterLogic_FilterLogicOr:
		return queryp.FilterLogicOr
	}
	return -1
}

func (fo FilterOp) FilterOp() queryp.FilterOp {
	switch fo {
	case FilterOp_FilterOpEquals:
		return queryp.FilterOpEquals
	case FilterOp_FilterOpNotEquals:
		return queryp.FilterOpNotEquals
	case FilterOp_FilterOpLessThan:
		return queryp.FilterOpLessThan
	case FilterOp_FilterOpLessThanEqual:
		return queryp.FilterOpLessThanEqual
	case FilterOp_FilterOpGreaterThan:
		return queryp.FilterOpGreaterThan
	case FilterOp_FilterOpGreaterThanEqual:
		return queryp.FilterOpGreaterThanEqual
	case FilterOp_FilterOpLike:
		return queryp.FilterOpLike
	case FilterOp_FilterOpNotLike:
		return queryp.FilterOpNotLike
	case FilterOp_FilterOpILike:
		return queryp.FilterOpILike
	case FilterOp_FilterOpNotILike:
		return queryp.FilterOpNotILike
	case FilterOp_FilterOpRegexp:
		return queryp.FilterOpRegexp
	case FilterOp_FilterOpNotRegexp:
		return queryp.FilterOpNotRegexp
	case FilterOp_FilterOpIRegexp:
		return queryp.FilterOpIRegexp
	case FilterOp_FilterOpNotIRegexp:
		return queryp.FilterOpNotIRegexp
	}
	return -1
}

func (ft FilterType) FilterType() queryp.FilterType {
	switch ft {
	case FilterType_FilterTypeNotFound:
		return queryp.FilterTypeNotFound
	case FilterType_FilterTypeSimple:
		return queryp.FilterTypeSimple
	case FilterType_FilterTypeString:
		return queryp.FilterTypeString
	case FilterType_FilterTypeNumeric:
		return queryp.FilterTypeNumeric
	case FilterType_FilterTypeTime:
		return queryp.FilterTypeTime
	case FilterType_FilterTypeBool:
		return queryp.FilterTypeBool
	}
	return -1
}

func SortQP(stpb []*SortTerm) queryp.Sort {
	if len(stpb) == 0 {
		return nil
	}
	s := make(queryp.Sort, 0, len(stpb))
	for _, st := range stpb {
		s = append(s, queryp.SortTerm{
			Field: st.Field,
			Desc:  st.Desc,
		})
	}
	return s
}
