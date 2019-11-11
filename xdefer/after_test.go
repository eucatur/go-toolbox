package xdefer

import "testing"

func TestAfter(t *testing.T) {
	var (
		name1 = "Henrique"
		name2 = "Pedro Henrique"
	)

	var strt = &struct {
		Name string
	}{
		Name: name1,
	}

	func() {
		after := New()
		defer after.Exec()
		after.Do(func() { strt.Name = name2 })
	}()

	if strt.Name != name1 {
		t.Error("The Do func was not executed.")
		return
	}

	func() {
		after := New()
		defer after.Exec()
		after.Do(func() { strt.Name = name1 })
		after.Cancel()
	}()

	if strt.Name == name1 {
		t.Error("The Do function should not be canceled.")
		return
	}
}
