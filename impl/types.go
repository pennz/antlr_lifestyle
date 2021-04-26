package impl

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

type ImplementationError struct {
	Msg string
}

func (e *ImplementationError) Error() string {
	return e.Msg
}
