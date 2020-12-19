package queryp

// Sort is a slice of SortTerms
type Sort []SortTerm

// SortFields is how you specify what fields are available to select for sorting
type SortFields map[string]Field

func (st SortTerm) String() string {

	if st.Desc {
		return "-" + st.Field
	}
	return st.Field

}
