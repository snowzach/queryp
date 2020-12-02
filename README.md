# QueryP

NOTE: This is under development and is not yet officially released. Expect breaking changes for a bit.

This is a Go library for generating (mostly) database agnostic query parameters with:
* Filter
* Sort
* Limit/Offsets

# Parsing
The library also supports parsing a string with a specific format into the query parameters. It's most
useful as a library for parsing GET query parameters.

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
	* !:~	- Not Regular Expression Match (case-insensitive

	* ()	- Parenthesis can be use for precedence.

# Thanks
Special thanks to my employer https://geneticnetworks.com for letting me work on this and open source it.

# TODO
* Allow filter types where the internal field name supports aliasing and possibly interfaces
* Rewrite parser to something a little more intelligent
