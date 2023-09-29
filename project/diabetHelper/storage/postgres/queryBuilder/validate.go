package queryBuilder

import "errors"

var operators = []string{And, Or}

func isValidOperator(operator string) bool {
	for i := range operators {
		if operator == operators[i] {
			return true
		}
	}
	return false
}

func (c *clauses) validateOrderBy() error {
	if c.isOffset && c.isLimit {
		return errors.New("'ORDER BY' must be before 'OFFSET' and 'LIMIT'")
	}
	if c.isOffset {
		return errors.New("'ORDER BY' must be before 'OFFSET'")
	}
	if c.isLimit {
		return errors.New("'ORDER BY' must be before 'LIMIT'")
	}

	if c.isOrderBy {
		return errors.New("'ORDER BY' is already specified")
	}

	return nil
}
