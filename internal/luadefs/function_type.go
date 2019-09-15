package luadefs

// FunctionType lets you determine if it's client-side, server-side or both
type FunctionType int8

const (
	// SharedFunctionType is a shared-function
	SharedFunctionType FunctionType = iota

	// ClientFunctionType is a client-side function
	ClientFunctionType

	// ServerFunctionType is a server-side function
	ServerFunctionType
)
