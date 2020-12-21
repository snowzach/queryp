package queryp

import "strings"

type FilterType int

const (
	FilterTypeNotFound FilterType = 0
	FilterTypeSimple   FilterType = 1
	FilterTypeString   FilterType = 2
	FilterTypeNumeric  FilterType = 3
	FilterTypeTime     FilterType = 4
	FilterTypeBool     FilterType = 5
)

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
