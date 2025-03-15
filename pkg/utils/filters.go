package utils

import (
	"fmt"
	"strings"
)

type Filter struct {
	Key      string
	Value    interface{}
	Operator FilterOperator
	AndOr    string
}

type FilterOperator int

const (
	FilterOperatorEqual FilterOperator = iota
	FilterOperatorLike
	FilterOperatorIsNull
	FilterOperatorBiggerOrEqual
	FilterOperatorSmallerOrEqual
)

func (f FilterOperator) String() string {
	return [...]string{"=", "LIKE", "IS NULL", ">=", "<="}[f]
}

// BuildWhereStatement builds a WHERE statement for a SQL query based on the filters provided for example "SELECT * FROM X %s LIMIT ?, ?"
func BuildWhereStatement(query *string, filters ...Filter) []interface{} {
	var whereStatement string
	var whereArgs []interface{}
	if len(filters) > 0 {
		whereStatement = "WHERE "
		for i, filter := range filters {
			if i > 0 {
				if filter.AndOr == "" || !strings.EqualFold(filter.AndOr, "or") {
					whereStatement += " AND "
				} else {
					whereStatement += " OR "
				}
			}
			if filter.Operator == FilterOperatorLike {
				whereStatement += fmt.Sprintf("%s LIKE ?", filter.Key)
				whereArgs = append(whereArgs, fmt.Sprintf("%%%s%%", filter.Value))
				continue
			}
			if filter.Value == nil {
				whereStatement += fmt.Sprintf("%s IS NULL", filter.Key)
				continue
			}

			if filter.Operator == FilterOperatorEqual {
				whereStatement += fmt.Sprintf("%s = ?", filter.Key)
				whereArgs = append(whereArgs, filter.Value)
				continue
			}

			if filter.Operator == FilterOperatorBiggerOrEqual {
				whereStatement += fmt.Sprintf("%s >= ?", filter.Key)
				whereArgs = append(whereArgs, filter.Value)
				continue
			}

			if filter.Operator == FilterOperatorSmallerOrEqual {
				whereStatement += fmt.Sprintf("%s <= ?", filter.Key)
				whereArgs = append(whereArgs, filter.Value)
				continue
			}
		}
	}
	*query = fmt.Sprintf(*query, whereStatement)
	return whereArgs
}
