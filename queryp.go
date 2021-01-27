package queryp

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// Field is just a string
type Field = string

type QueryParameters struct {
	Filter  Filter  `json:"filter"`
	Sort    Sort    `json:"sort"`
	Options Options `json:"options"`
	Limit   int64   `json:"limit"`
	Offset  int64   `json:"offset"`
}

// Handles parsing query requests with complex matching and precedence
var (
	optionParser = regexp.MustCompile("^option\\[(.*)\\]$")
)

func ParseRawQuery(rq string) (*QueryParameters, error) {
	q, err := url.PathUnescape(rq)
	if err != nil {
		return nil, err
	}
	return ParseQuery(q)
}

// ParseQuery converts a string into query parameters
// This loosely follows standard HTTP URL encoding
func ParseQuery(q string) (*QueryParameters, error) {

	qp := &QueryParameters{
		Sort:    make(Sort, 0),
		Options: make(Options),
	}

	if len(q) == 0 {
		return qp, nil
	}

	var pos int
	var err error

	var parse func(depth int) (Filter, error)
	parse = func(depth int) (Filter, error) {

		start := true
		filter := make([]FilterTerm, 0)

		for {

			// Are we done parsing this loop
			if pos == len(q) || (q[pos] == ')' && depth > 0) {
				break
			}

			// Parse the filter logic
			var logic FilterLogic // Default is START if omitted
			if start {
				logic = FilterLogicAnd
				start = false
			} else {
				var found bool
				logic, found = FilterLogicSymToFilterLogic[string(q[pos])]
				if !found {
					return nil, fmt.Errorf("invalid filter logic at pos %d", pos)
				}
				pos++
			}

			if pos == len(q) {
				return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
			}

			// Is there a sub-filter?
			if q[pos] == '(' {
				// Eat paren
				pos++
				if pos == len(q) {
					return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
				}
				// Parse the sub-filter
				subFilter, err := parse(depth + 1)
				if err != nil {
					return nil, err
				}
				if pos == len(q) {
					return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
				}
				// Ensure paren closed
				if q[pos] != ')' {
					return nil, fmt.Errorf("missing close parenthesis at pos: %d", pos)
				}
				pos++
				// Append sub-filter
				filter = append(filter, FilterTerm{
					Logic:     logic,
					SubFilter: subFilter, // Parse, handle redundant parens
				})
				if pos == len(q) { // nothing else left to parse
					return filter, nil
				}

			} else {

				// Field identifier
				var fieldb strings.Builder
				if pos == len(q) {
					return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
				}
				if q[pos] == '"' {
					pos++
				fieldLoopQuote:
					for {
						if pos == len(q) {
							return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
						}
						switch q[pos] {
						case '"':
							pos++
							break fieldLoopQuote
						case '\\': // Escape whatever
							pos++
							if pos == len(q) {
								return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
							}
							fieldb.WriteByte(q[pos])
						default:
							fieldb.WriteByte(q[pos])
						}
						pos++
					}
				} else {
				fieldLoop:
					for {
						if pos == len(q) {
							return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
						}
						switch q[pos] {
						case '!', '=', '~', ':', '<', '>', '@':
							break fieldLoop
						default:
							fieldb.WriteByte(q[pos])
						}
						pos++
					}

				}
				if fieldb.Len() == 0 {
					return nil, fmt.Errorf("invalid field name at pos %d", pos)
				}
				field := fieldb.String()

				// Operator
				var opb strings.Builder
			opLoop:
				for {
					if pos == len(q) {
						return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
					}
					switch q[pos] {
					case '!', '=', '~', ':', '<', '>', '@':
						opb.WriteByte(q[pos])
					default:
						break opLoop
					}
					pos++
				}

				op, found := FilterOpSymToFilterOp[opb.String()]
				if !found {
					return nil, fmt.Errorf("invalid filter operation: %s at pos %d", opb.String(), pos)
				}

				// value identifier
				var valueQuoted bool
				var valueb strings.Builder
				if pos == len(q) {
					return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
				}

				// Support field=(value1,value2,"value3") in syntax
				if q[pos] == '(' {
					pos++
					values := make([]interface{}, 0)
					for {
						var multaValueQuoted bool
						if pos == len(q) {
							return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
						}
						if q[pos] == '"' {
							pos++
							multaValueQuoted = true
						multiValueLoopQuote:
							for {
								if pos == len(q) {
									return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
								}
								switch q[pos] {
								case '"':
									pos++
									break multiValueLoopQuote
								case '\\': // Escape whatever
									pos++
									if pos == len(q) {
										return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
									}
									valueb.WriteByte(q[pos])
								default:
									valueb.WriteByte(q[pos])
								}
								pos++
							}
						} else {
						multiValueLoop:
							for {
								if pos == len(q) {
									// End of input
									break multiValueLoop
								}
								switch q[pos] {
								case ',', ')':
									break multiValueLoop
								default:
									valueb.WriteByte(q[pos])
								}
								pos++
							}
						}
						if valueb.Len() == 0 {
							return nil, fmt.Errorf("invalid value at pos %d", pos)
						}
						// If the value was null and not quoted, we want it to be an actual nil
						if value := valueb.String(); value == NullValue && !multaValueQuoted {
							values = append(values, nil)
						} else {
							values = append(values, value)
						}
						valueb.Reset()

						// See if there are more values or done with values
						if pos == len(q) {
							return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
						}
						if q[pos] == ')' {
							pos++
							break // Done paring options
						} else if q[pos] == ',' { // additional values
							pos++
						} else {
							return nil, fmt.Errorf("unexpected value at pos %d", pos)
						}
					}

					filter = append(filter, FilterTerm{
						Logic: logic,
						Op:    op,
						Field: Field(field),
						Value: values,
					})

					continue

					// Single quoted value
				} else if q[pos] == '"' {
					pos++
					valueQuoted = true
				valueLoopQuote:
					for {
						if pos == len(q) {
							return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
						}
						switch q[pos] {
						case '"':
							pos++
							break valueLoopQuote
						case '\\': // Escape whatever
							pos++
							if pos == len(q) {
								return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
							}
							valueb.WriteByte(q[pos])
						default:
							valueb.WriteByte(q[pos])
						}
						pos++
					}

					// Single unquoted value
				} else {
				valueLoop:
					for {
						if pos == len(q) {
							// End of input
							break valueLoop
						}
						switch q[pos] {
						case '&', '|', ')':
							break valueLoop
						default:
							valueb.WriteByte(q[pos])
						}
						pos++
					}

				}

				if valueb.Len() == 0 {
					return nil, fmt.Errorf("invalid value at pos %d", pos)
				}
				value := valueb.String()

				// We will only handle options at depth 0
				if depth == 0 {
					switch field {
					case "limit":
						if op != FilterOpEquals {
							return nil, fmt.Errorf("invalid operation for limit option")
						}
						qp.Limit, err = strconv.ParseInt(value, 10, 64)
						if err != nil {
							return nil, fmt.Errorf("invalid value %s for limit", value)
						}
						field = "" // mark field parsed
					case "offset":
						if op != FilterOpEquals {
							return nil, fmt.Errorf("invalid operation for offset option")
						}
						qp.Offset, err = strconv.ParseInt(value, 10, 64)
						if err != nil {
							return nil, fmt.Errorf("invalid value %s for offset", value)
						}
						field = "" // mark field parsed
					case "sort":
						if op != FilterOpEquals {
							return nil, fmt.Errorf("invalid operation for sort option")
						}
						for _, sortField := range strings.Split(value, ",") {
							if len(sortField) > 1 && sortField[0] == '-' { // Reverse sort
								qp.Sort = append(qp.Sort, SortTerm{Field: sortField[1:], Desc: true})
							} else {
								qp.Sort = append(qp.Sort, SortTerm{Field: sortField})
							}
						}
						field = "" // mark field parsed
					case "option":
						if op != FilterOpEquals {
							return nil, fmt.Errorf("invalid operation for option")
						}
						if qp.Options == nil {
							qp.Options = make(Options)
						}
						// Is it option[key]=value format?
						if m := optionParser.FindStringSubmatch(field); len(m) > 0 {
							qp.Options[m[1]] = value
						} else {
							for _, optionField := range strings.Split(value, ",") {
								qp.Options[optionField] = "true" // Default value when not explicitly set
							}
						}
						field = "" // mark field parsed
					}
				}

				// This is a field/value option
				if len(field) > 0 { // field will be empty if it was parsed as an option
					// Handle an actual nil value
					if value == NullValue && !valueQuoted {
						filter = append(filter, FilterTerm{
							Logic: logic,
							Op:    op,
							Field: Field(field),
							Value: nil,
						})
					} else {
						filter = append(filter, FilterTerm{
							Logic: logic,
							Op:    op,
							Field: Field(field),
							Value: value,
						})
					}
				}
			}
		}

		return filter, nil
	}

	// Parse the filter
	qp.Filter, err = parse(0)
	if err != nil {
		return nil, err
	}

	return qp, nil

}

// Reset clears query parameters for reuse
func (qp *QueryParameters) Reset() {
	qp.Filter = nil
	qp.Sort = nil
	qp.Options = nil
	qp.Limit = 0
	qp.Offset = 0
}

// String will turn query parameters back into a string that is equivalent to what can be parsed
func (qp *QueryParameters) String() string {

	var sb strings.Builder

	// Filter
	sb.WriteString(Filter(qp.Filter).String())

	// The Limit and Offset
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

	// Options
	if options := qp.Options.String(); len(options) > 0 {
		if sb.Len() > 0 {
			sb.WriteString("&")
		}
		sb.WriteString(options)
	}

	// Sort
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

// SafeField will quote and escape fields if they contain op symbols or quotes
func SafeField(field string) string {

	if strings.ContainsAny(field, FieldNeedsQuote) {
		return `"` + strings.ReplaceAll(field, `"`, `\"`) + `"`
	}
	return field

}

// Convert our values to a string
func ValueString(value interface{}) string {

	var sval string
	switch s := value.(type) {
	case string:
		sval = s
	case []byte:
		sval = string(s)
	case fmt.Stringer:
		sval = s.String()
	case bool:
		return strconv.FormatBool(s)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case int32:
		return strconv.FormatInt(int64(s), 10)
	case int64:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatUint(uint64(s), 10)
	case uint8:
		return strconv.FormatUint(uint64(s), 10)
	case uint16:
		return strconv.FormatUint(uint64(s), 10)
	case uint32:
		return strconv.FormatUint(uint64(s), 10)
	case uint64:
		return strconv.FormatUint(uint64(s), 10)
	case nil:
		return "null"
	case []interface{}:
		values := make([]string, len(s), len(s))
		for i, v := range s {
			values[i] = ValueString(v)
		}
		return "(" + strings.Join(values, ",") + ")"
	default:
		sval = fmt.Sprintf("%v", value)
	}

	if strings.ContainsAny(sval, `&|)"`) {
		return `"` + strings.ReplaceAll(sval, `"`, `\"`) + `"`
	} else if sval == NullValue {
		return `"null"`
	}

	return sval

}
