package queryp

import (
	"bytes"
	"encoding/json"
)

type QueryParameters struct {
	Filter  Filter  `json:"filter"`
	Sort    Sort    `json:"sort"`
	Options Options `json:"options"`
	Limit   int     `json:"limit"`
	Offset  int     `json:"offset"`
}

type Filter []FilterTerm

type FilterTerm struct {
	Logic FilterLogic `json:"logic"`
	Op    FilterOp    `json:"op"`
	Field Field       `json:"field"`
	Value string      `json:"value"`

	SubFilter Filter `json:"sub_filter,omitempty"`
}

type Field = string // Alias

type SortFields []string

type Sort []SortTerm

type SortTerm struct {
	Field Field `json:"field"`
	Desc  bool  `json:"desc"`
}

type Options map[string]string // Just a lookup of string

func (o Options) HasOption(option string) bool {
	_, found := o[option]
	return found
}

func (o Options) MarshalJSON() ([]byte, error) {
	ret := make([]string, 0, len(o))
	for key := range o {
		ret = append(ret, key)
	}
	return json.Marshal(ret)
}

func (qp *QueryParameters) String() string {
	b := &bytes.Buffer{}
	e := json.NewEncoder(b)
	e.SetEscapeHTML(false)
	if err := e.Encode(qp); err != nil {
		panic(err)
	}
	return b.String()
}

func (qp *QueryParameters) PrettyString() string {

	b := &bytes.Buffer{}
	e := json.NewEncoder(b)
	e.SetEscapeHTML(false)
	e.SetIndent("", "  ")
	if err := e.Encode(qp); err != nil {
		panic(err)
	}
	return b.String()

}
