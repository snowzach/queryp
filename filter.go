package queryp

import (
	"fmt"
	"strconv"
	"strings"
)

// Filter is a slice of FilterTerms that are combined to make the full filter
type Filter []FilterTerm

// Field is just a string
type Field = string

// FilterField is an interface used to allow specifying a filter for a particular filter type.
// if GetFieldName returns an empty string, it will assume the FieldName is the same as the field name specified in the map
type FilterField interface {
	GetFieldName() string
	GetFilterType() FilterType
}

type FilterFieldTypes map[Field]FilterField

// FindFilterType search the FilterFieldTypes to find a matching field name.
// First it searches for exact matches and then
func (fft FilterFieldTypes) FindFilterType(search string) (string, FilterType) {

	if filterField, found := fft[search]; found {
		if name := filterField.GetFieldName(); name != "" {
			return name, filterField.GetFilterType()
		}
		return search, filterField.GetFilterType()
	}

	// Search for the suffix in the list of filters
	search = "." + search
	for filter, filterField := range fft {
		if strings.HasSuffix(filter, search) {
			if name := filterField.GetFieldName(); name != "" {
				return name, filterField.GetFilterType()
			}
			return filter, filterField.GetFilterType()
		}
	}
	return "", FilterTypeNotFound

}

// GetFieldName returns empty field indicating it should use the same name as the filter type
func (ft FilterType) GetFieldName() Field {
	return ""
}

// GetFilterType for the default filter types
func (ft FilterType) GetFilterType() FilterType {
	return ft
}

// FilterFieldCustom is used to have a different name for a field from what can be specified in the parsable string
type FilterFieldCustom struct {
	FieldName  string
	FilterType FilterType
}

// GetFieldName to return filter name from a Custom filter type
func (ffc FilterFieldCustom) GetFieldName() Field {
	return ffc.FieldName
}

// GetFilterType to return the filter type from a Custom filter type
func (ffc FilterFieldCustom) GetFilterType() FilterType {
	return ffc.FilterType
}

// Convienience function for appending to a filter
func (f Filter) Append(logic FilterLogic, field Field, op FilterOp, value interface{}) Filter {

	var sval string
	switch s := value.(type) {
	case string:
		sval = s
	case []byte:
		sval = string(s)
	case fmt.Stringer:
		sval = s.String()
	case bool:
		sval = strconv.FormatBool(s)
	case float64:
		sval = strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		sval = strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		sval = strconv.FormatInt(int64(s), 10)
	case int8:
		sval = strconv.FormatInt(int64(s), 10)
	case int32:
		sval = strconv.FormatInt(int64(s), 10)
	case int64:
		sval = strconv.FormatInt(int64(s), 10)
	case uint:
		sval = strconv.FormatUint(uint64(s), 10)
	case uint8:
		sval = strconv.FormatUint(uint64(s), 10)
	case uint16:
		sval = strconv.FormatUint(uint64(s), 10)
	case uint32:
		sval = strconv.FormatUint(uint64(s), 10)
	case uint64:
		sval = strconv.FormatUint(uint64(s), 10)
	case nil:
		sval = "null"
	default:
		sval = fmt.Sprintf("%v", value)
	}
	return append(f, FilterTerm{Logic: logic, Field: field, Op: op, Value: sval})

}

// Filter turns it into it's string representations
func (f Filter) String() string {

	var sb strings.Builder

	for i, ft := range f {
		if i > 0 {
			sb.WriteString(ft.Logic.String())
		}
		if len(ft.SubFilter) > 0 {
			_, _ = sb.WriteString("(")
			_, _ = sb.WriteString(Filter(ft.SubFilter).String())
			_, _ = sb.WriteString(")")
		} else {
			sb.WriteString(ft.String())
		}
	}

	return sb.String()

}

// String converts FilterTerm into it's string representation
func (ft FilterTerm) String() string {
	return SafeField(ft.Field) + ft.Op.String() + SafeValue(ft.Value)
}
