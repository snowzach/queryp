package queryp

import "strings"

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
