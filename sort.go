package queryp

import "strings"

type Sort []*SortTerm

type SortFields []SortField

type SortField string

func (sf SortField) GetSortName() string {
	m := strings.SplitN(string(sf), ":", 2)
	return m[0]
}

func (sf SortField) GetFieldName() string {
	m := strings.SplitN(string(sf), ":", 2)
	if len(m) > 1 {
		return m[1]
	}
	return m[0]
}
