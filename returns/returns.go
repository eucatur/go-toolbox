package returns

func All(v ...interface{}) []interface{} {
	return v
}

func First(v ...interface{}) interface{} {
	return v[0]
}

func Second(v ...interface{}) interface{} {
	return v[1]
}
