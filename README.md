# QueryP

[![GoDoc](https://godoc.org/github.com/snowzach/queryp?status.svg)](https://godoc.org/github.com/snowzach/queryp)

This is a Go library for generating (mostly) database agnostic query parameters with:
* Filter
* Sort
* Limit/Offsets

# Parsing
The library also supports parsing a string with a specific format into the query parameters. It's most
useful as a library for parsing GET-like query parameters.

For example: 
```
field1=value1&(field2=value2|field3=value3)&limit=10&offset=10&sort=name,-description

Get all things where field1=value1 AND ( field2=value2 OR field3=value3 ) sort by name ascending, description descending limit to 10 records and skip the first 10
```

## Supported Operators:
    * &		- Logic AND (as well as splitting operators)
    * |		- Logic OR

	* =     - Exactly Equals
	* !=    - Not Equals
	* < 	- Less Than
	* >		- Greater Than
	* <=	- Less than or equals
	* >=	- Greater than or equals
	* =~	- Like (supports % wildcards)
	* !=~   - Not Like (supports % wildcards)
	* =~~   - Like case-insensitive (supports % wildcards)
	* !=~~	- Not Like case-insensitive (supports % wildcards)
	* :		- Regular Expression Match
	* !:	- Not Regular Expression Match
	* :~	- Regular Expression Match (case-insensitive
	* !:~	- Not Regular Expression Match (case-insensitive)
	* @     - Bits are set (for numeric values)
	* @~    - Bits are clear (for numeric values)

	* ()	- Parenthesis can be use for precedence.

# Drivers
Database drivers are built into the package
* qppg - Postgres syntax driver

# Usage

```
queryString := "field1=value1&(field2=value2|field3=value3)&limit=10&offset=10&sort=name,-description"

qp, err := queryp.ParseQuery(queryString)
if err != nil {
	log.Fatal(err)
}

var queryClause strings.Builder
var queryParams = []interface{}{}

// Fields we are allowed to filter by
filterFields := queryp.FilterFieldTypes{
	"table_name.thing_id":          queryp.FilterTypeSimple, // Simple matching good for equal/not equal
	"table_name.thing_name":        queryp.FilterTypeString, // Matching with like and regexp
	"table_name.created":           queryp.FilterTypeTime,   // Matching based on times
	"table_name.exists":            queryp.FilterTypeBool,   // Supports true/false only
	"table_name.quantity":          queryp.FilterTypeNumeric,// Supports numeric comparisons
	"someotherfield":               queryp.FilterFieldCustom{FieldName: "whoathird", FilterType: queryp.FilterTypeString}, // Use one field name in URL and another for database field
}

// Fields we are allowed to sort by
sortFields := queryp.SortFields{
	"source.source_id",
	"source.source_name",
}

// Default sort if none specified
if len(qp.Sort) == 0 {
	qp.Sort = queryp.Sort{queryp.SortTerm{Field: "source.source_name", Desc: false}}
}

// If we will filter
if len(qp.Filter) > 0 {
	queryClause.WriteString(" WHERE ")
}

// Generate the WHERE criteria for postgres
if err := qppg.FilterQuery(filterFields, qp.Filter, &queryClause, &queryParams); err != nil {
	return nil, 0, err
}

// Generate the sort criteria at the end of the query clause 
if err := qppg.SortQuery(sortFields, qp.Sort, &queryClause, &queryParams); err != nil {
	return nil, 0, err
}

// Append limit and offset if specified
if qp.Limit > 0 {
	queryClause.WriteString(" LIMIT " + strconv.Itoa(qp.Limit))
}
if qp.Offset > 0 {
	queryClause.WriteString(" OFFSET " + strconv.Itoa(qp.Offset))
}

query, err := db.Query("SELECT * FROM table_name "+queryClause.String(), queryParams...)
if err != nil {
	log.Fatal(err)
}

```

# Thanks
Special thanks to my employer https://geneticnetworks.com for letting me work on this and open source it.


