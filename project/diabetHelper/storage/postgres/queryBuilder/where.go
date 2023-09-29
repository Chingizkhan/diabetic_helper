package queryBuilder

import (
	"fmt"
	"strconv"
)

const (
	And   string = "AND"
	Or    string = "OR"
	where string = "WHERE"
)

const invalidOperator string = "queryBuilder.%s invalid operator '%s'"

// Where - sets the 'WHERE' clause using the custom operator
// example: WHERE name = 'John'
func (qb *QueryBuilder) Where(column string, parameter interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "Where()", operator)
	}

	qb.updateQueryBuilder(
		column+" $"+strconv.Itoa(qb.counter)+"\n",
		parameter,
		operator,
	)

	return nil
}

// WhereEqual - sets the 'WHERE' clause using the '=' operator
// example: WHERE name = 'John'
func (qb *QueryBuilder) WhereEqual(column string, parameter interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "WhereEqual()", operator)
	}

	qb.updateQueryBuilder(
		column+" = $"+strconv.Itoa(qb.counter)+"\n",
		parameter,
		operator,
	)

	return nil
}

// WhereNotEqual - sets the 'WHERE' clause using the '!=' operator
// example: WHERE name != 'John'
func (qb *QueryBuilder) WhereNotEqual(column string, parameter interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "WhereNotEqual()", operator)
	}

	qb.updateQueryBuilder(
		column+" != $"+strconv.Itoa(qb.counter)+"\n",
		parameter,
		operator,
	)

	return nil
}

// WhereGreater - sets the 'WHERE' clause using the '>' operator
// example: WHERE id > '5'
func (qb *QueryBuilder) WhereGreater(column string, parameter interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "WhereGreater()", operator)
	}

	qb.updateQueryBuilder(
		column+" > $"+strconv.Itoa(qb.counter)+"\n",
		parameter,
		operator,
	)

	return nil
}

// WhereLess - sets the 'WHERE' clause using the '<' operator
// example: WHERE id < '5'
func (qb *QueryBuilder) WhereLess(column string, parameter interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "WhereLess()", operator)
	}

	qb.updateQueryBuilder(
		column+" < $"+strconv.Itoa(qb.counter)+"\n",
		parameter,
		operator,
	)

	return nil
}

// WhereGreaterOrEqual - sets the 'WHERE' clause using the '>=' operator
// example: WHERE id >= '5'
func (qb *QueryBuilder) WhereGreaterOrEqual(column string, parameter interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "WhereGreaterOrEqual()", operator)
	}

	qb.updateQueryBuilder(
		column+" >= $"+strconv.Itoa(qb.counter)+"\n",
		parameter,
		operator,
	)

	return nil
}

// WhereLessOrEqual - sets the 'WHERE' clause using the '<=' operator
// example: WHERE id <= '5'
func (qb *QueryBuilder) WhereLessOrEqual(column string, parameter interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "WhereLessOrEqual()", operator)
	}

	qb.updateQueryBuilder(
		column+" <= $"+strconv.Itoa(qb.counter)+"\n",
		parameter,
		operator,
	)

	return nil
}

// WhereStartsWith - sets the 'WHERE' clause using the 'LIKE' operator
// example: WHERE name LIKE 'John%'
func (qb *QueryBuilder) WhereStartsWith(column string, parameter interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "WhereStartsWith()", operator)
	}

	qb.updateQueryBuilder(
		column+" LIKE $"+strconv.Itoa(qb.counter)+"\n",
		fmt.Sprint(parameter)+"%",
		operator,
	)

	return nil
}

// WhereEndsWith - sets the 'WHERE' clause using the 'LIKE' operator
// example: WHERE name LIKE '%Doe'
func (qb *QueryBuilder) WhereEndsWith(column string, parameter interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "WhereEndsWith()", operator)
	}

	qb.updateQueryBuilder(
		column+" LIKE $"+strconv.Itoa(qb.counter)+"\n",
		"%"+fmt.Sprint(parameter),
		operator,
	)

	return nil
}

// WhereStartsAndEnds - sets the 'WHERE' clause using the 'LIKE' operator
// example: WHERE name LIKE '%John%'
func (qb *QueryBuilder) WhereStartsAndEnds(column string, parameter interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "WhereAny()", operator)
	}

	qb.updateQueryBuilder(
		column+" LIKE $"+strconv.Itoa(qb.counter)+"\n",
		"%"+fmt.Sprint(parameter)+"%",
		operator,
	)

	return nil
}

// WhereIn - sets the 'WHERE' clause using the 'IN' operator
// example: WHERE name IN ('Bob', 'John', 'Dean')
func (qb *QueryBuilder) WhereIn(column string, parameters []interface{}, operator string) error {
	if !isValidOperator(operator) {
		return fmt.Errorf(invalidOperator, "WhereAny()", operator)
	}

	var s string
	paramLen := len(parameters)
	for i, v := range parameters {
		s += "'" + fmt.Sprint(v) + "'"
		if i+1 != paramLen {
			s += ", "
		}
	}
	qb.counter--
	qb.updateQueryBuilder(
		column+" IN ("+s+")\n",
		nil,
		operator,
	)

	return nil
}

func (qb *QueryBuilder) updateQueryBuilder(s string, parameter interface{}, operator string) {
	if qb.counter != 1 {
		qb.operator = operator
	}

	qb.query += qb.operator + " " + s
	if parameter != nil {
		qb.parameters = append(qb.parameters, parameter)
	}
	qb.counter++
}
