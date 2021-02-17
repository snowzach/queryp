package queryp

import (
	"strings"
)

type Options map[string]string // Just a lookup of string

func (o Options) Get(option string) string {
	if o == nil {
		return ""
	}
	if value, found := o[option]; found {
		return value
	}
	return ""
}

func (o Options) Has(option string) bool {
	if o == nil {
		return false
	}
	_, found := o[option]
	return found
}

func (o *Options) Set(option string, value string) Options {
	if *o == nil {
		*o = make(Options)
	}
	(*o)[option] = value
	return *o
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
