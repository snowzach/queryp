package queryp

import (
	"bytes"
	"encoding/json"
)

type Filter []FilterTerm
type Field = string // Alias
type SortFields []string
type Sort []SortTerm

type Options map[string]string // Just a lookup of string

func (o Options) HasOption(option string) bool {
	_, found := o[option]
	return found
}

func (o Options) Set(option string, value string) Options {
	if o == nil {
		o = make(Options)
	}
	o[option] = value
	return o
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
