// Package check é utilizado para simplificar o uso de condicionais (if) no go.
// Esse pacote funciona da mesma forma que o operador ternário em outras linguagens.
package check

// If se condition for um valor verdadeiro, o valor trueValue será retornado.
// Se condition for um valor falso, o valor falseValue será retornado.
//
//  resultado := check.If(2 > 1, "true", "false").(string) //true
//
func If(condition bool, trueValue interface{}, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

// IfFunc se condition for um valor verdadeiro, a função trueFunc será executada.
// Se condition for um valor falso, a função falseFunc será executada.
//
//  trueFunc := func() interface{} { return "true" }
//  falseFunc := func() interface{} { return "false" }
//  resultado := check.If(2 < 1, trueFunc, falseFunc).(string)  //false
//
func IfFunc(condition bool, trueFunc func() interface{}, falseFunc func() interface{}) interface{} {
	if condition {
		return trueFunc()
	}
	return falseFunc()
}
