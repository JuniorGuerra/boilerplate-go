package kernel

type (
	Interceptor struct{}

	HandlerFunc func(ctx AppContext)
)
