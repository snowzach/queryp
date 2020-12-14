package queryp

import (
	"fmt"
	"strconv"
	"strings"
)

type FilterField interface {
	GetFieldName() string
	GetFilterType() FilterType
}

type FilterFieldTypes map[Field]FilterField

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

func (ft FilterType) GetFieldName() Field {
	return ""
}

func (ft FilterType) GetFilterType() FilterType {
	return ft
}

type FilterFieldCustom struct {
	FieldName  string
	FilterType FilterType
}

func (ffc FilterFieldCustom) GetFieldName() Field {
	return ffc.FieldName
}

func (ffc FilterFieldCustom) GetFilterType() FilterType {
	return ffc.FilterType
}

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
