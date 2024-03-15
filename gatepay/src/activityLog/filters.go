package activityLog

type Filters []*Filter

type Filter struct {
	Property           string `json:"property"`
	Value              string `json:"value"`
	RelationalOperator string `json:"relational_operator"`
	LogicalOperator    string `json:"logical_operator"`
}

func (f *Filter) query() string {
	return f.Property + " " + f.RelationalOperator + " ? "
}

func FilterByUserEmail(email string) *Filter {
	return &Filter{
		Property:           "author",
		Value:              email,
		RelationalOperator: "=",
		LogicalOperator:    "and",
	}
}
