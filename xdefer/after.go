package xdefer

// After is a struct
type After struct {
	tasks    []func()
	canceled bool
}

// New starts a new instance of struct After.
func New() After {
	return After{}
}

// Exec performs all functions added by the Do function.
func (after *After) Exec() {
	if after.canceled {
		return
	}

	for _, task := range after.tasks {
		task()
	}
}

// Do adds a function to perform later.
func (after *After) Do(task func()) {
	after.tasks = append(after.tasks, task)
}

// Cancel the execution of functions added by the Do function.
func (after *After) Cancel() {
	after.canceled = true
}
