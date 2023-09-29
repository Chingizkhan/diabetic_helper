package queryBuilder

import (
	"strconv"
)

const (
	ascending  string = "ASC"
	descending string = "DESC"
)

// Offset sets the 'OFFSET' clause
// example: OFFSET 5
func (qb *QueryBuilder) Offset(offset uint32) {
	if offset == 0 {
		return
	}
	qb.query += "OFFSET" + " " + strconv.Itoa(int(offset)) + "\n"
	qb.clauses.isOffset = true
}

// Limit sets the 'LIMIT' clause
// example: LIMIT 10
func (qb *QueryBuilder) Limit(limit uint32) {
	if limit == 0 {
		return
	}
	qb.query += "LIMIT" + " " + strconv.Itoa(int(limit)) + "\n"
	qb.clauses.isLimit = true
}

// OrderBy sets the 'ORDER BY' clause
// example: ORDER BY created_at DESC
func (qb *QueryBuilder) OrderBy(column string, isDescending bool) error {
	err := qb.clauses.validateOrderBy()
	if err != nil {
		return err
	}

	sortStatement := getSortStatement(isDescending)
	qb.query += "ORDER BY" + " " + column + " " + sortStatement + "\n"
	qb.clauses.isOrderBy = true
	return nil
}

func (qb *QueryBuilder) GroupBy(column string) {
	qb.query += "GROUP BY" + " " + column + "\n"
	qb.clauses.isGroupBy = true
}

func getSortStatement(isDescending bool) string {
	orderByParameter := ascending
	if isDescending {
		orderByParameter = descending
	}
	return orderByParameter
}
