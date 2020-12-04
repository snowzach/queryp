package queryp

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// Handles parsing query requests with complex matching and precedence
var (
	optionParser = regexp.MustCompile("^option\\[(.*)\\]$")

	ErrCouldNotParse = errors.New("could not parse")
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

			// Parse the filter logic
			var logic FilterLogic // Default is START if omitted
			if start {
				logic = FilterLogicStart
				start = false
			} else {
				switch q[pos] {
				case '&':
					logic = FilterLogicAnd
				case '|':
					logic = FilterLogicOr
				default:
					return nil, fmt.Errorf("invalid filter logic at pos %d", pos)
				}
				pos++
			}

			if pos == len(q) {
				return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
			}

			// Is there a sub-filter?
			if q[pos] == '(' {
				log.Println("SUB")
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
						case '!', '=', '~', ':', '<', '>':
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
					case '!', '=', '~', ':', '<', '>':
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
				var valueb strings.Builder
				if pos == len(q) {
					return nil, fmt.Errorf("unexpected end of input at pos %d", pos)
				}
				if q[pos] == '"' {
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
						qp.Limit, err = strconv.Atoi(value)
						if err != nil {
							return nil, fmt.Errorf("invalid value %s for limit", value)
						}
						field = "" // mark field parsed
					case "offset":
						if op != FilterOpEquals {
							return nil, fmt.Errorf("invalid operation for offset option")
						}
						qp.Offset, err = strconv.Atoi(value)
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

				// Otherwise add a filter option
				if len(field) > 0 { // field will be empty if it was parsed as an option
					filter = append(filter, FilterTerm{
						Logic: logic,
						Op:    op,
						Field: Field(field),
						Value: value,
					})
				}
			}

			if pos == len(q) || (q[pos] == ')' && depth > 0) {
				break
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
