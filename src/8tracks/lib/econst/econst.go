package econst

// ContextKey represents the key id of the context
type ContextKey int

const (
	// RequestID is the context key name to hold the generated request id of incoming request
	RequestID ContextKey = iota
	// UserID is the user id of the incoming request
	UserID
)
