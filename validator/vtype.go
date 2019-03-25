package validator

type vType struct {
	validations map[string]vFunc
}

func (v *vType) addValidation(n string, vf vFunc) {
	v.validations[n] = vf
}

func (v *vType) getValidation(n string) (vFunc, bool) {
	vf, ok := v.validations[n]
	return vf, ok
}
