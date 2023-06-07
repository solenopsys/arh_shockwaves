package compilers

type CompileExecutor interface {
	Compile(params map[string]string) error
}

type CompileCommand struct {
	LibName      string
	LibDirectory string
}
