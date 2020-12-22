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
			case queryp.FilterLogicAnd:
				queryClause.WriteString(" AND ")
			case queryp.FilterLogicOr:
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
			if filterType == queryp.FilterTypeNotFound {
				return fmt.Errorf("could not find field: %s", ft.Field)
			}
			queryClause.WriteString(field)

			switch filterType {

			case queryp.FilterTypeSimple, queryp.FilterTypeNumeric, queryp.FilterTypeTime:

				switch ft.Op {
				case queryp.FilterOpEquals:
					queryClause.WriteString(" = ")
				case queryp.FilterOpNotEquals:
					queryClause.WriteString(" != ")
				case queryp.FilterOpLessThan:
					queryClause.WriteString(" < ")
				case queryp.FilterOpLessThanEqual:
					queryClause.WriteString(" <= ")
				case queryp.FilterOpGreaterThan:
					queryClause.WriteString(" > ")
				case queryp.FilterOpGreaterThanEqual:
					queryClause.WriteString(" >= ")
				case queryp.FilterOpIn:
					queryClause.WriteString(" && ") // Overlap operator
				default:
					return fmt.Errorf("invalid op %s for field %s", ft.Op.String(), field)
				}

				*queryParams = append(*queryParams, ft.Value)
				queryClause.WriteString("$" + strconv.Itoa(len(*queryParams)))

			case queryp.FilterTypeString:
				switch ft.Op {
				case queryp.FilterOpEquals:
					queryClause.WriteString(" = ")
				case queryp.FilterOpNotEquals:
					queryClause.WriteString(" != ")
				case queryp.FilterOpLessThan:
					queryClause.WriteString(" < ")
				case queryp.FilterOpLessThanEqual:
					queryClause.WriteString(" <= ")
				case queryp.FilterOpGreaterThan:
					queryClause.WriteString(" > ")
				case queryp.FilterOpGreaterThanEqual:
					queryClause.WriteString(" >= ")
				case queryp.FilterOpLike:
					queryClause.WriteString(" LIKE ")
				case queryp.FilterOpNotLike:
					queryClause.WriteString(" NOT LIKE ")
				case queryp.FilterOpILike:
					queryClause.WriteString(" ILIKE ")
				case queryp.FilterOpNotILike:
					queryClause.WriteString(" NOT ILIKE ")
				case queryp.FilterOpRegexp:
					queryClause.WriteString(" ~ ")
				case queryp.FilterOpNotRegexp:
					queryClause.WriteString(" !~ ")
				case queryp.FilterOpIRegexp:
					queryClause.WriteString(" ~* ")
				case queryp.FilterOpNotIRegexp:
					queryClause.WriteString(" !~* ")
				case queryp.FilterOpIn:
					queryClause.WriteString(" && ") // Overlap operator
				default:
					return fmt.Errorf("invalid op %s for field %s", ft.Op.String(), field)
				}

				*queryParams = append(*queryParams, ft.Value)
				queryClause.WriteString("$" + strconv.Itoa(len(*queryParams)))

			case queryp.FilterTypeBool:

				switch ft.Op {
				case queryp.FilterOpEquals:
					queryClause.WriteString(" = ")
				case queryp.FilterOpNotEquals:
					queryClause.WriteString(" != ")
				case queryp.FilterOpIn:
					queryClause.WriteString(" && ") // Overlap operator
				default:
					return fmt.Errorf("invalid op %s for field %s", ft.Op.String(), field)
				}

				boolVal, err := strconv.ParseBool(queryp.ValueString(ft.Value))
				if err != nil {
					return fmt.Errorf("invalid bool value %s for field %s", ft.Value, field)
				}

				*queryParams = append(*queryParams, boolVal)
				queryClause.WriteString("$" + strconv.Itoa(len(*queryParams)))

			default:
				return fmt.Errorf("invalid filter type for field %s", field)

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
		var sortName, sortField string
		for sortName, sortField = range sortFields {
			if sortTerm.Field == sortName {
				found = true
				break
			}
		}
		// Check for a matching suffix
		if !found {
			sortTermFieldSuffix := "." + sortTerm.Field
			for sortName, sortField = range sortFields {
				if strings.HasSuffix(sortName, sortTermFieldSuffix) {
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
			if len(sortField) > 0 {
				queryClause.WriteString(sortField)
			} else {
				queryClause.WriteString(sortName)
			}
			if sortTerm.Desc {
				queryClause.WriteString(" DESC")
			}
		}
	}
	return nil

}
