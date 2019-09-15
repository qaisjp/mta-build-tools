package luadefs

// LuaDef represents a defined Lua definition
type LuaDef struct {
	FunctionName string
	Type         FunctionType

	// Extensions:
	// Class string
	// SourcePath string
	// Documentation string
}

// OnClient tells you if this definition is available client-side
func (d LuaDef) OnClient() bool {
	return d.Type != ServerFunctionType
}

// OnServer tells you if this definition is available server-side
func (d LuaDef) OnServer() bool {
	return d.Type != ClientFunctionType
}

// IsShared tells you if this definition is available both client-side and server-side
func (d LuaDef) IsShared() bool {
	return d.Type == SharedFunctionType
}
