package qppb

import (
	"github.com/snowzach/queryp"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func QueryParametersPB(qp *queryp.QueryParameters) *QueryParameters {

	return &QueryParameters{
		Filter:  FilterPB(qp.Filter),
		Sort:    SortPB(qp.Sort),
		Options: qp.Options,
		Limit:   qp.Limit,
		Offset:  qp.Offset,
	}

}

func FilterPB(ftqp []*queryp.FilterTerm) []*FilterTerm {
	if len(ftqp) == 0 {
		return nil
	}
	f := make([]*FilterTerm, 0, len(ftqp))
	for _, ft := range ftqp {
		value, _ := structpb.NewValue(ft.Value)
		f = append(f, &FilterTerm{
			Logic:     FilterLogicPB(ft.Logic),
			Field:     ft.Field,
			Op:        FilterOpPB(ft.Op),
			Value:     value,
			SubFilter: FilterPB(ft.SubFilter),
		})
	}
	return f
}

func FilterLogicPB(fl queryp.FilterLogic) FilterLogic {
	switch fl {
	case queryp.FilterLogicAnd:
		return FilterLogic_FilterLogicAnd
	case queryp.FilterLogicOr:
		return FilterLogic_FilterLogicOr
	}
	return -1
}

func FilterOpPB(fo queryp.FilterOp) FilterOp {
	switch fo {
	case queryp.FilterOpEquals:
		return FilterOp_FilterOpEquals
	case queryp.FilterOpNotEquals:
		return FilterOp_FilterOpNotEquals
	case queryp.FilterOpLessThan:
		return FilterOp_FilterOpLessThan
	case queryp.FilterOpLessThanEqual:
		return FilterOp_FilterOpLessThanEqual
	case queryp.FilterOpGreaterThan:
		return FilterOp_FilterOpGreaterThan
	case queryp.FilterOpGreaterThanEqual:
		return FilterOp_FilterOpGreaterThanEqual
	case queryp.FilterOpLike:
		return FilterOp_FilterOpLike
	case queryp.FilterOpNotLike:
		return FilterOp_FilterOpNotLike
	case queryp.FilterOpILike:
		return FilterOp_FilterOpILike
	case queryp.FilterOpNotILike:
		return FilterOp_FilterOpNotILike
	case queryp.FilterOpRegexp:
		return FilterOp_FilterOpRegexp
	case queryp.FilterOpNotRegexp:
		return FilterOp_FilterOpNotRegexp
	case queryp.FilterOpIRegexp:
		return FilterOp_FilterOpIRegexp
	case queryp.FilterOpNotIRegexp:
		return FilterOp_FilterOpNotIRegexp
	}
	return -1
}

func FilterTypePB(ft queryp.FilterType) FilterType {
	switch ft {
	case queryp.FilterTypeNotFound:
		return FilterType_FilterTypeNotFound
	case queryp.FilterTypeSimple:
		return FilterType_FilterTypeSimple
	case queryp.FilterTypeString:
		return FilterType_FilterTypeString
	case queryp.FilterTypeNumeric:
		return FilterType_FilterTypeNumeric
	case queryp.FilterTypeTime:
		return FilterType_FilterTypeTime
	case queryp.FilterTypeBool:
		return FilterType_FilterTypeBool
	}
	return -1
}

func SortPB(stqp queryp.Sort) []*SortTerm {
	if len(stqp) == 0 {
		return nil
	}
	s := make([]*SortTerm, 0, len(stqp))
	for _, st := range stqp {
		s = append(s, &SortTerm{
			Field: st.Field,
			Desc:  st.Desc,
		})
	}
	return s
}
