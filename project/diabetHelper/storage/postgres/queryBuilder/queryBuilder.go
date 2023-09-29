package queryBuilder

type QueryBuilder struct {
	Table      string
	columns    string
	query      string
	operator   string
	counter    int
	clauses    clauses
	parameters []interface{}
}

type clauses struct {
	isOrderBy bool
	isGroupBy bool
	isOffset  bool
	isLimit   bool
}

func NewQueryBuilder(table string, columns string, as string) *QueryBuilder {
	qb := &QueryBuilder{
		Table:      table,
		columns:    columns,
		operator:   where,
		counter:    1,
		parameters: []interface{}{},
	}
	switch as {
	case "":
		qb.query = "SELECT" + " " + columns + "\nFROM" + " " + table + "\n"
	default:
		qb.query = "SELECT" + " " + columns + "\nFROM" + " " + table + " " + as + "\n"
	}

	return qb
}

func (qb *QueryBuilder) InnerJoin(table string, as string, on string) {
	qb.query += "INNER JOIN " + table + " AS " + as + " ON " + on + "\n"
}

// GetQuery returns a query
func (qb *QueryBuilder) GetQuery() string {
	return qb.query
}

// GetParameters returns parameters
func (qb *QueryBuilder) GetParameters() []interface{} {
	return qb.parameters
}
