package query

import "fmt"

type Query struct {
	query string
	args  []interface{}
}

func New(startment string, args ...interface{}) *Query {
	return &Query{query: startment, args: args}
}

func (q Query) String() string {
	fmt.Println(q.query)
	return q.query
}

func (q Query) Args() []interface{} {
	return q.args
}

func (q *Query) And(condition bool, value string, arg interface{}) *Query {
	if condition {
		q.query += value
		q.args = append(q.args, arg)
	}
	return q
}
