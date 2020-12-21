package queryp

// Sort is a slice of SortTerms
type Sort []SortTerm

type SortTerm struct {
	Field Field `json:"field"`
	Desc  bool  `json:"desc"`
}

// SortFields is how you specify what fields are available to select for sorting
type SortFields map[string]Field

func (st SortTerm) String() string {

	if st.Desc {
		return "-" + st.Field
	}
	return st.Field

}
