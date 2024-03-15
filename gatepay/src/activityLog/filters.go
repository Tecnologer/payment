package activityLog

type Filters []*Filter

type Filter struct {
	Field              string `json:"field"`
	Value              string `json:"value"`
	RelationalOperator string `json:"relational_operator"`
	LogicalOperator    string `json:"logical_operator"`
}

func (f *Filter) query() string {
	return f.Field + " " + f.RelationalOperator + " ? "
}

func FilterByUserEmail(email string) *Filter {
	return &Filter{
		Field:              "author",
		Value:              email,
		RelationalOperator: "=",
		LogicalOperator:    "and",
	}
}
