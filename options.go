package queryp

import (
	"sort"
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

func (o *Options) Set(option string, value string) {
	if *o == nil {
		*o = make(Options)
	}
	(*o)[option] = value
}

func (o Options) String() string {

	if o == nil || len(o) == 0 {
		return ""
	}

	// Sort keys to have consistent order
	keys := make([]string, 0, len(o))
	for key := range o {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for i, key := range keys {
		if i > 0 {
			sb.WriteString("&")
		}
		value := o[key]
		if value == "true" {
			sb.WriteString("option=" + key)
		} else {
			sb.WriteString("option[" + key + "]=" + value)
		}
	}
	return sb.String()

}
