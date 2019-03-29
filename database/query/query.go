package query

// Query is the main structure of the query
type Query struct {
	query string
	args  []interface{}
}

// New start a new query
func New(startment string, args ...interface{}) *Query {
	return &Query{query: startment, args: args}
}

// Add concatenates the value in the query and add the args
func (q *Query) Add(value string, args ...interface{}) *Query {
	q.query += value
	q.args = append(q.args, args...)
	return q
}

// AddIf concatenates the value in the query and add the args if the condition is true
func (q *Query) AddIf(condition bool, value string, args ...interface{}) *Query {
	if condition {
		q.Add(value, args...)
	}
	return q
}

// String returns query as string
func (q Query) String() string {
	return q.query
}

// Args returns all args
func (q Query) Args() []interface{} {
	return q.args
}
