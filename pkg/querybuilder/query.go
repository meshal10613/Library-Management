package querybuilder

type QueryParams struct {
	Page   int
	Limit  int
	Search string
	SortBy string
	Order  string
}

func NewQueryParams() QueryParams {
	return QueryParams{
		Page:  1,
		Limit: 10,
	}
}
