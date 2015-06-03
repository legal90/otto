package infrastructure

// Mock is a mock implementation of the Infrastructure interface.
type Mock struct {
	CompileCalled  bool
	CompileContext *Context
	CompileResult  *CompileResult
	CompileErr     error
}

func (m *Mock) Compile(ctx *Context) (*CompileResult, error) {
	m.CompileCalled = true
	m.CompileContext = ctx
	return m.CompileResult, m.CompileErr
}

func (m *Mock) Flavors() []string {
	return nil
}