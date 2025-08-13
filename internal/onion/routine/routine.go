package routine

type Routine interface {
}

type routineImpl struct{}

func NewRoutine() Routine {
	return &routineImpl{}
}
