package queryp

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	FilterTypeNotFound FilterType = iota
	FilterTypeSimple
	FilterTypeString
	FilterTypeNumeric
	FilterTypeTime
	FilterTypeBool
)

type FilterType int

type FilterField interface {
	Name() string
	Type() FilterType
}

type FilterFieldTypes map[Field]FilterField

func (fft FilterFieldTypes) FindFilterType(search string) (string, FilterType) {

	if filterField, found := fft[search]; found {
		if name := filterField.Name(); name != "" {
			return name, filterField.Type()
		}
		return search, filterField.Type()
	}

	// Search for the suffix in the list of filters
	search = "." + search
	for filter, filterField := range fft {
		if strings.HasSuffix(filter, search) {
			if name := filterField.Name(); name != "" {
				return name, filterField.Type()
			}
			return filter, filterField.Type()
		}
	}
	return "", FilterTypeNotFound

}

func (ft FilterType) Name() string {
	return ""
}

func (ft FilterType) Type() FilterType {
	return ft
}

type FilterFieldCustom struct {
	FieldName  string
	FilterType FilterType
}

func (ffc FilterFieldCustom) Name() string {
	return ffc.FieldName
}

func (ffc FilterFieldCustom) Type() FilterType {
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
