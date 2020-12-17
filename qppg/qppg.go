package qppg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/snowzach/queryp"
)

// FilterQuery will update the queryClause and queryParams with filter values
func FilterQuery(fft queryp.FilterFieldTypes, filter queryp.Filter, queryClause *strings.Builder, queryParams *[]interface{}) error {

	for i, ft := range filter {

		if i > 0 {
			switch ft.Logic {
			case queryp.FilterLogic_And:
				queryClause.WriteString(" AND ")
			case queryp.FilterLogic_Or:
				queryClause.WriteString(" OR ")
			}
		}

		if ft.SubFilter != nil {
			queryClause.WriteString("(")
			if err := FilterQuery(fft, ft.SubFilter, queryClause, queryParams); err != nil {
				return err
			}
			queryClause.WriteString(")")
		} else {

			// Lookup the filter
			field, filterType := fft.FindFilterType(ft.Field)
			if filterType == queryp.FilterType_NotFound {
				return fmt.Errorf("could not find field: %s", ft.Field)
			}
			queryClause.WriteString(field)

			switch filterType {

			case queryp.FilterType_Simple, queryp.FilterType_Numeric, queryp.FilterType_Time:

				switch ft.Op {
				case queryp.FilterOp_Equals:
					queryClause.WriteString(" = ")
				case queryp.FilterOp_NotEquals:
					queryClause.WriteString(" != ")
				case queryp.FilterOp_LessThan:
					queryClause.WriteString(" < ")
				case queryp.FilterOp_LessThanEqual:
					queryClause.WriteString(" <= ")
				case queryp.FilterOp_GreaterThan:
					queryClause.WriteString(" > ")
				case queryp.FilterOp_GreaterThanEqual:
					queryClause.WriteString(" >= ")
				default:
					return fmt.Errorf("invalid op %s for field %s", ft.Op.String(), field)
				}

				*queryParams = append(*queryParams, ft.Value)
				queryClause.WriteString("$" + strconv.Itoa(len(*queryParams)))

			case queryp.FilterType_String:
				switch ft.Op {
				case queryp.FilterOp_Equals:
					queryClause.WriteString(" = ")
				case queryp.FilterOp_NotEquals:
					queryClause.WriteString(" != ")
				case queryp.FilterOp_LessThan:
					queryClause.WriteString(" < ")
				case queryp.FilterOp_LessThanEqual:
					queryClause.WriteString(" <= ")
				case queryp.FilterOp_GreaterThan:
					queryClause.WriteString(" > ")
				case queryp.FilterOp_GreaterThanEqual:
					queryClause.WriteString(" >= ")
				case queryp.FilterOp_Like:
					queryClause.WriteString(" LIKE ")
				case queryp.FilterOp_NotLike:
					queryClause.WriteString(" NOT LIKE ")
				case queryp.FilterOp_ILike:
					queryClause.WriteString(" ILIKE ")
				case queryp.FilterOp_NotILike:
					queryClause.WriteString(" NOT ILIKE ")
				case queryp.FilterOp_Regexp:
					queryClause.WriteString(" ~ ")
				case queryp.FilterOp_NotRegexp:
					queryClause.WriteString(" !~ ")
				case queryp.FilterOp_IRegexp:
					queryClause.WriteString(" ~* ")
				case queryp.FilterOp_NotIRegexp:
					queryClause.WriteString(" !~* ")
				default:
					return fmt.Errorf("invalid op %s for field %s", ft.Op.String(), field)
				}

				*queryParams = append(*queryParams, ft.Value)
				queryClause.WriteString("$" + strconv.Itoa(len(*queryParams)))

			case queryp.FilterType_Bool:

				switch ft.Op {
				case queryp.FilterOp_Equals:
					queryClause.WriteString(" = ")
				case queryp.FilterOp_NotEquals:
					queryClause.WriteString(" != ")
				default:
					return fmt.Errorf("invalid op %s for field %s", ft.Op.String(), field)
				}

				boolVal, err := strconv.ParseBool(ft.Value)
				if err != nil {
					return fmt.Errorf("invalid bool value %s for field %s", ft.Value, field)
				}

				*queryParams = append(*queryParams, boolVal)
				queryClause.WriteString("$" + strconv.Itoa(len(*queryParams)))

			}
		}
	}

	return nil
}

// SortQuery will update the queryClause and queryParams with sort values
func SortQuery(sortFields queryp.SortFields, sort queryp.Sort, queryClause *strings.Builder, queryParams *[]interface{}) error {

	if len(sort) == 0 {
		return nil
	}

	var first = true
	for _, sortTerm := range sort {
		// Search for exact match
		var found bool
		var sortField queryp.SortField
		for _, sortField = range sortFields {
			if sortTerm.Field == sortField.GetSortName() {
				found = true
				break
			}
		}
		// Check for a matching suffix
		if !found {
			sortTermFieldSuffix := "." + sortTerm.Field
			for _, sortField = range sortFields {
				if strings.HasSuffix(sortField.GetSortName(), sortTermFieldSuffix) {
					found = true
					break
				}
			}
		}
		// Found a field, build the order by clause
		if found {
			if first {
				queryClause.WriteString(" ORDER BY ")
				first = false
			} else {
				queryClause.WriteString(", ")
			}
			queryClause.WriteString(sortField.GetFieldName())
			if sortTerm.Desc {
				queryClause.WriteString(" DESC")
			}
		}
	}
	return nil

}
