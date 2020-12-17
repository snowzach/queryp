package queryp

import (
	"bytes"
	"encoding/json"
)

type Filter []*FilterTerm
type Field = string // Alias

type Options map[string]string // Just a lookup of string

func (qp *QueryParameters) HasOption(option string) bool {
	_, found := qp.Options[option]
	return found
}

func (qp *QueryParameters) SetOption(option string, value string) {
	if qp.Options == nil {
		qp.Options = make(map[string]string)
	}
	qp.Options[option] = value
}

// func (qp *QueryParameters) String() string {
// 	b := &bytes.Buffer{}
// 	e := json.NewEncoder(b)
// 	e.SetEscapeHTML(false)
// 	if err := e.Encode(qp); err != nil {
// 		panic(err)
// 	}
// 	return b.String()
// }

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
