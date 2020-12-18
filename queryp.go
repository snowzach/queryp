package queryp

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

type Filter []FilterTerm
type Field = string // Alias

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

func (o Options) String() string {

	if o == nil {
		return ""
	}

	var sb strings.Builder
	var join bool
	for k, v := range o {
		if join {
			sb.WriteString("&")
		}
		if v == "true" {
			sb.WriteString("option=" + k)
		} else {
			sb.WriteString("option[" + k + "]=" + v)
		}
		join = true
	}

	return sb.String()
}

func (qp *QueryParameters) String() string {

	var sb strings.Builder

	sb.WriteString(Filter(qp.Filter).String())

	if qp.Limit > 0 {
		if sb.Len() > 0 {
			sb.WriteString("&")
		}
		sb.WriteString("limit=" + strconv.FormatInt(qp.Limit, 10))
	}

	if qp.Offset > 0 {
		if sb.Len() > 0 {
			sb.WriteString("&")
		}
		sb.WriteString("offset=" + strconv.FormatInt(qp.Offset, 10))
	}

	if options := qp.Options.String(); len(options) > 0 {
		if sb.Len() > 0 {
			sb.WriteString("&")
		}
		sb.WriteString(options)
	}

	if len(qp.Sort) > 0 {
		sb.WriteString("&sort=")
		for i, st := range qp.Sort {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(st.String())
		}
	}

	return sb.String()

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
