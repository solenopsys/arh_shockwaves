package internal

type CompileExecutor interface {
	Compile(params map[string]string) error
}

type CompileParamsExtractor interface {
	Extract(name string, path string) map[string]string
}

type CompileCommand struct {
	LibName      string
	LibDirectory string
}
